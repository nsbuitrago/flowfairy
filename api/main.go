package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
  "github.com/nsbuitrago/flowfairy/api/models/FlowData"
)

func main() {
  //router setup
  router := chi.NewRouter()
  router.Use(middleware.Logger)
  router.Use(cors.Handler(cors.Options{
  AllowedOrigins: []string{"http://localhost:3000"},
  AllowedMethods: []string{"GET", "POST"},
  }))


  //routes
  router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Operational"))
  })

  router.Post("/load", LoadFlow)
  http.ListenAndServe(":8000", router)
}

func LoadFlow(w http.ResponseWriter, r *http.Request) {
  
  responseObject := &FlowData{}
  r.ParseMultipartForm(32 >> 20) // 32 mb max memory

}

