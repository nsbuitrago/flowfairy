package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nsbuitrago/fcs2/fcs2"
	"github.com/nsbuitrago/flowfairy/models"
	"net/http"
  "sync"
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
	router.Get("/status", GetStatus)
	router.Post("/load", LoadData)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  err := json.NewEncoder().Encode("Operational")
  if err != nil {
    return
  }
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 >> 20) // 32 mb max memory (rest will be in disk)
	if err != nil {
		return
	}

  var wg sync.WaitGroup
  var response models.FlowData
  filenames := []string{"fcs_file_0"}

  for {
    wg.Add(1)
    go func() {
      defer wg.Done()
      filename := filenames[0] // todo: get filename
      filenames = append(filenames[:0], filenames[0:]...) // is this the best way to do this?
      _, fileHeader, err := r.FormFile(filename)
      if err != nil {
        panic(err) // this and subsequent panics should instead form an err response to send back
      }

      file, err := fileHeader.Open()
      if err != nil {
        panic(err)
      }

      fcsMetaData, fcsData, err := fcs2.NewDecoder(file).Decode()
      err = file.Close() 
      if err != nil {
        panic(err)
        }

      response.ID = 1
      response.Events = fcsData
      response.MetaData = fcsMetaData
    }()

  wg.Wait()
	//render.Render(w, r, fr)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

