package models

import (
  "net/http"
  "fmt"

)

// implement FlowData response structure

type FlowData struct {
  ID int `json:"id"` // internal ID
  Events [][]float64 `json:"events"` 
  MetaData map[string]string `json:"metadata"` 
}

func (f *FlowData) Bind(r *http.Request) error {
  if f.ID == 0 {
    return fmt.Errorf("ID is a required field and must be nonzero")
  }

  return nil
}

func (*FlowData) Render(w http.ResponseWriter, r *http.Request) error {
  return nil
}
