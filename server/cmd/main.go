package main

import (
	"context"
	"embed"
	"github.com/c0nscience/your-turn-to-roll/pkg/fight"
	"github.com/c0nscience/your-turn-to-roll/pkg/public"
	"github.com/c0nscience/your-turn-to-roll/pkg/session"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

type spaHandler struct {
	staticFS   embed.FS
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.staticPath, path)

	_, err = h.staticFS.Open(path)
	if os.IsNotExist(err) {
		index, err := h.staticFS.ReadFile(filepath.Join(h.staticPath, h.indexPath))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statics, err := fs.Sub(h.staticFS, h.staticPath)
	http.FileServer(http.FS(statics)).ServeHTTP(w, r)
}

func main() {
	log.Println("starting server")

	r := mux.NewRouter()

	r.HandleFunc("/api/session/{id}/continue", session.Continue).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/session/start", session.Start).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/fight/start", fight.Start).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/fight/{id}/update", fight.Update).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/fight/{id}/ws", fight.Sync)

	if public.IsEmbedded {
		spa := spaHandler{
			staticFS:   public.Box,
			staticPath: "dist",
			indexPath:  "index.html",
		}
		r.PathPrefix("/").Handler(spa)
	}

	var handler http.Handler = r
	if !public.IsEmbedded {
		corsOpts := cors.New(cors.Options{
			AllowedOrigins: []string{
				"*",
			},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodHead,
			},
			AllowedHeaders: []string{"*"},
		})
		handler = corsOpts.Handler(r)
	}

	srv := &http.Server{
		Addr:         ":8081",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.LoggingHandler(os.Stdout, handler),
	}

	go func() {
		log.Println("server started on 8081")

		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
