// Package server teaches wiring a store to HTTP handlers (create + list).
package server

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// Store is a concurrency-safe in-memory ticket list.
type Store struct {
	mu    sync.Mutex
	items []Ticket
	next  int
}

func (s *Store) Add(t Ticket) Ticket {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	t.ID = s.next
	s.items = append(s.items, t)
	return t
}

func (s *Store) Get(id int) (Ticket, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range s.items {
		if t.ID == id {
			return t, true
		}
	}
	return Ticket{}, false
}

func (s *Store) Update(id int, t Ticket) (Ticket, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.items {
		if s.items[i].ID == id {
			t.ID = id
			s.items[i] = t
			return t, true
		}
	}
	return Ticket{}, false
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.items {
		if s.items[i].ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return true
		}
	}
	return false
}

// List returns a copy of every stored ticket.
func (s *Store) List() []Ticket {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Ticket, len(s.items))
	copy(out, s.items)
	return out
}

// Server serves tickets over HTTP.
type Server struct {
	store *Store
	mux   *http.ServeMux
}

// NewServer builds a Server with the ticket collection routes registered.
func NewServer() *Server {
	s := &Server{store: &Store{}, mux: http.NewServeMux()}
	s.mux.HandleFunc("POST /tickets", s.handleCreate)
	s.mux.HandleFunc("GET /tickets", s.handleList)
	return s
}

// Handler exposes the routes.
func (s *Server) Handler() http.Handler { return s.mux }

// handleCreate decodes a ticket, validates it, and answers 201 with the
// stored ticket as JSON (400 for bad input).
func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var t Ticket
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}
	if t.Title == "" {
		http.Error(w, "title required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s.store.Add(t))
}

// handleList answers 200 with every stored ticket as JSON.
func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.store.List())
}
