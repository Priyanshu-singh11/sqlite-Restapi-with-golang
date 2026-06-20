package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Priyanshu-singh11/Restapi/internal/config"
	"github.com/Priyanshu-singh11/Restapi/internal/sqlite"
	"github.com/Priyanshu-singh11/Restapi/internal/student"
)

// webFiles embeds the frontend at compile time, so serving it no longer
// depends on the working directory you launch the binary from. The "web"
// folder must sit in the same directory as this main.go file.
//
//go:embed web
//go:embed web
var webFiles embed.FS

// withCORS lets the frontend call the API even if it's ever served from a
// different origin (a separate dev server, a different port, etc). When the
// frontend is served by this same process — the default setup below — these
// headers are simply unused, but they're safe to leave on. If you deploy
// this publicly, replace "*" with your real frontend origin.
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	//load config
	cfg := config.MustLoad()

	//DB connection
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Db initialise", slog.String("env", cfg.Env))

	//routing
	router := http.NewServeMux()

	// API
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateStudentById(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteStudentById(storage))

	// Frontend — serves the embedded web/index.html (and any other static
	// assets in web/) at the site root. Embedding means this works no
	// matter what directory you run the binary from.
	webContent, err := fs.Sub(webFiles, "web")
	if err != nil {
		log.Fatal(err)
	}
	router.Handle("/", http.FileServer(http.FS(webContent)))

	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: withCORS(router),
	}

	fmt.Printf("Server running on http://%s\n", cfg.HTTPServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(
		done,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error(
			"failed to shutdown server",
			slog.String("error", err.Error()),
		)
	}

	slog.Info("server stopped")
}
