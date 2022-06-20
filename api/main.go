package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nsbuitrago/fcs2/fcs2"
	"github.com/nsbuitrago/flowfairy/models"
	"net/http"
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
		_, err := w.Write([]byte("Operational"))
		if err != nil {
			return
		}
	})

	router.Post("/load", LoadFlow)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}

func LoadFlow(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 >> 20) // 32 mb max memory (rest will be in disk)
	if err != nil {
		return
	}

	_, fileHeader, err := r.FormFile("fcs_file_0")
	file, err := fileHeader.Open()
	fcsMetaData, fcsData, err := fcs2.NewDecoder(file).Decode()
	err = file.Close()
	if err != nil {
		return
	}

	testResponse := models.FlowData{
		ID:     1,
		Events: fcsData,
		MetaData: fcsMetaData,
	}

	//render.Render(w, r, fr)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&testResponse)
	if err != nil {
		return
	}

}
