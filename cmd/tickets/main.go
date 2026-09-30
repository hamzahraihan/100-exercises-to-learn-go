// Command tickets runs the course capstone as a real server: a JSON ticket
// API with file persistence, environment config, and graceful shutdown.
//
//	PORT=8080 DATA_FILE=tickets.json go run ./cmd/tickets
//
// It is deliberately self-contained (standard library only, no exercise
// packages) so it runs on a fresh clone whether or not the exercises are
// solved. Every pattern below was taught in the course: environment config,
// atomic saves, boot loading, mutex-guarded storage, JSON handlers with
// precise statuses, and draining shutdown.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Config holds server settings. Same rules as the config lesson: every
// setting has a default, empty means unset, and a bad PORT falls back
// instead of crashing the server.
type Config struct {
	Port     int
	DataFile string
}

// ConfigFromEnv reads PORT (default 8080) and DATA_FILE (default
// tickets.json) from the environment.
func ConfigFromEnv() Config {
	port := 8080
	if v := os.Getenv("PORT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 1 && n <= 65535 {
			port = n
		}
	}
	dataFile := os.Getenv("DATA_FILE")
	if dataFile == "" {
		dataFile = "tickets.json"
	}
	return Config{Port: port, DataFile: dataFile}
}

// Ticket is a support request.
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status"`
}

// Store is a concurrency-safe durable ticket list: mutations hold the
// write seam across in-memory change + file save, reads use RLock.
type Store struct {
	mu    sync.RWMutex
	items []Ticket
	next  int
	data  string // state file path; empty means in-memory only (tests)
}

// NewStore builds a Store preloaded with ts, continuing ids past the max.
func NewStore(ts []Ticket, dataFile string) *Store {
	s := &Store{data: dataFile}
	for _, t := range ts {
		if t.ID > s.next {
			s.next = t.ID
		}
		s.items = append(s.items, t)
	}
	return s
}

// saveLocked persists current items; caller must hold write lock.
func (s *Store) saveLocked() error {
	if s.data == "" {
		return nil
	}
	out := make([]Ticket, len(s.items))
	copy(out, s.items)
	return saveTickets(s.data, out)
}

// Add stores t with a fresh id.
func (s *Store) Add(t Ticket) (Ticket, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	t.ID = s.next
	s.items = append(s.items, t)
	if err := s.saveLocked(); err != nil {
		return t, err
	}
	return t, nil
}

// Get finds a ticket by id.
func (s *Store) Get(id int) (Ticket, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, t := range s.items {
		if t.ID == id {
			return t, true
		}
	}
	return Ticket{}, false
}

// Update replaces the ticket with id, preserving the path id.
func (s *Store) Update(id int, t Ticket) (Ticket, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.items {
		if s.items[i].ID == id {
			t.ID = id
			s.items[i] = t
			if err := s.saveLocked(); err != nil {
				return t, true, err
			}
			return t, true, nil
		}
	}
	return Ticket{}, false, nil
}

// Delete removes the ticket with id.
func (s *Store) Delete(id int) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.items {
		if s.items[i].ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			if err := s.saveLocked(); err != nil {
				return true, err
			}
			return true, nil
		}
	}
	return false, nil
}

// List returns a copy of every stored ticket.
func (s *Store) List() []Ticket {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Ticket, len(s.items))
	copy(out, s.items)
	return out
}

// saveTickets writes ts atomically: temp file in the target's directory,
// then rename, so a crash never leaves a half-written state file.
func saveTickets(path string, ts []Ticket) error {
	out, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), "tickets-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	failed := true
	defer func() {
		if failed {
			os.Remove(tmpName)
		}
	}()
	if _, err = tmp.Write(out); err != nil {
		tmp.Close()
		return err
	}
	if err = tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err = tmp.Close(); err != nil {
		return err
	}
	failed = false
	return os.Rename(tmpName, path)
}

// loadTickets reads boot state: missing file means a fresh start (nil,
// nil); anything else propagates, wrapped with the path.
func loadTickets(path string) ([]Ticket, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var ts []Ticket
	if err := json.Unmarshal(data, &ts); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}
	return ts, nil
}

// Server serves tickets over HTTP; durability lives in Store.
type Server struct {
	store *Store
	mux   *http.ServeMux
}

// NewServer wires routes to handlers sharing one store.
func NewServer(store *Store) *Server {
	s := &Server{store: store, mux: http.NewServeMux()}
	s.mux.HandleFunc("POST /tickets", s.handleCreate)
	s.mux.HandleFunc("GET /tickets", s.handleList)
	s.mux.HandleFunc("GET /tickets/{id}", s.handleGet)
	s.mux.HandleFunc("PUT /tickets/{id}", s.handleUpdate)
	s.mux.HandleFunc("DELETE /tickets/{id}", s.handleDelete)
	return s
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: %v", err)
	}
}

func decodeBody(w http.ResponseWriter, r *http.Request, v any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return false
	}
	return true
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var t Ticket
	if !decodeBody(w, r, &t) {
		return
	}
	if t.Title == "" {
		http.Error(w, "title required", http.StatusBadRequest)
		return
	}
	t, err := s.store.Add(t)
	if err != nil {
		http.Error(w, "save failed", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusCreated, t)
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.store.List())
}

func (s *Server) parseID(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("id"))
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := s.parseID(r)
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

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := s.parseID(r)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	var t Ticket
	if !decodeBody(w, r, &t) {
		return
	}
	if t.Title == "" {
		http.Error(w, "title required", http.StatusBadRequest)
		return
	}
	updated, ok, err := s.store.Update(id, t)
	if err != nil {
		http.Error(w, "save failed", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := s.parseID(r)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	ok, err := s.store.Delete(id)
	if err != nil {
		http.Error(w, "save failed", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func newHTTPServer(cfg Config, srv *Server) *http.Server {
	return &http.Server{
		Addr:              ":" + strconv.Itoa(cfg.Port),
		Handler:           srv.mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func waitAndShutdown(ctx context.Context, httpSrv *http.Server, timeout time.Duration) error {
	<-ctx.Done()
	log.Print("shutting down, draining in-flight requests")
	shutCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return httpSrv.Shutdown(shutCtx)
}

func main() {
	cfg := ConfigFromEnv()

	ts, err := loadTickets(cfg.DataFile)
	if err != nil {
		log.Fatalf("load %s: %v", cfg.DataFile, err)
	}
	srv := NewServer(NewStore(ts, cfg.DataFile))
	httpSrv := newHTTPServer(cfg, srv)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("tickets listening on :%d (state: %s)", cfg.Port, cfg.DataFile)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("serve: %v", err)
		}
	}()

	if err := waitAndShutdown(ctx, httpSrv, 5*time.Second); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
	log.Print("stopped cleanly")
}
