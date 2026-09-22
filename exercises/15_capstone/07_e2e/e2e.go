// Package e2e teaches wiring the whole ticket API: every handler works,
// but no route is registered yet.
package e2e

import (
	"encoding/json"
	"net/http"
	"strconv"
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

// NewServer builds the server. Handlers and store already work.
func NewServer() *Server {
	s := &Server{store: &Store{}, mux: http.NewServeMux()}
	s.mux.HandleFunc("POST /tickets", s.handleCreate)
	s.mux.HandleFunc("GET /tickets", s.handleList)
	s.mux.HandleFunc("GET /tickets/{id}", s.handleGet)
	s.mux.HandleFunc("PUT /tickets/{id}", s.handleUpdate)
	s.mux.HandleFunc("DELETE /tickets/{id}", s.handleDelete)
	return s
}

// Handler exposes the routes.
func (s *Server) Handler() http.Handler { return s.mux }

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

// handleCreate decodes a ticket, validates it, and answers 201.
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
	writeJSON(w, http.StatusCreated, s.store.Add(t))
}

// handleList answers 200 with every stored ticket.
func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.List())
}

// handleGet answers 200 with one ticket (400 bad id, 404 missing).
func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	t, ok := s.store.Get(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// handleUpdate replaces one ticket (400 bad input, 404 missing).
func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	var t Ticket
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}
	if t.Title == "" {
		http.Error(w, "title required", http.StatusBadRequest)
		return
	}
	updated, ok := s.store.Update(id, t)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// handleDelete removes one ticket (400 bad id, 404 missing, else 204).
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	if !s.store.Delete(id) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
