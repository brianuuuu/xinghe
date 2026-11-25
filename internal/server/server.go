// Package server exposes the local HTTP API and serves the Web workbench.
package server

import (
	"encoding/json"
	"net/http"
	"time"
)

// Config configures the Xinghe HTTP server.
type Config struct {
	Addr    string
	DataDir string
}

// Server wraps core HTTP handlers.
type Server struct {
	cfg Config
	mux *http.ServeMux
}

// New creates a new Server instance with minimal baseline routes.
func New(cfg Config) *Server {
	s := &Server{
		cfg: cfg,
		mux: http.NewServeMux(),
	}
	s.routes()
	return s
}

// Handler returns the root http.Handler.
func (s *Server) Handler() http.Handler {
	return s.mux
}

func (s *Server) routes() {
	s.mux.HandleFunc("/api/health", s.handleHealth)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status":   "ok",
		"service":  "xinghe",
		"data_dir": s.cfg.DataDir,
		"time":     time.Now().Format(time.RFC3339),
	})
}
