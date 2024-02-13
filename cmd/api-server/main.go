package main

import (
	"net/http"

	"github.com/samvimes01/go-rss/internal/api/middleware"
	"github.com/samvimes01/go-rss/internal/api/routes"
	"github.com/samvimes01/go-rss/internal/env"
)


func main() {
	envir := env.NewEnv()
	cfg := routes.NewAPIConfig(envir)

	mux := http.NewServeMux()

	routes.Setup(envir, mux, cfg)
	
	corsHandler := middleware.Cors(mux)
	
	server := http.Server{
		Handler: corsHandler,
		Addr:    envir.Host + ":" + envir.Port,
	}

	server.ListenAndServe()
}
