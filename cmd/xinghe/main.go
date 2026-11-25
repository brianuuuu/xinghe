package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/brianuuuu/xinghe/internal/server"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8787", "HTTP listen address (default loopback)")
	dataDir := flag.String("data", ".xinghe", "Runtime data directory (Git docs, SQLite, drafts, objects, logs)")
	flag.Parse()

	srv := server.New(server.Config{
		Addr:    *addr,
		DataDir: *dataDir,
	})

	log.Printf("xinghe server listening on http://%s (data=%s)", *addr, *dataDir)
	if err := http.ListenAndServe(*addr, srv.Handler()); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
