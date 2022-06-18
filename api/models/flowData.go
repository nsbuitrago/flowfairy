package models

import "github.com/nsbuitrago/fcs2/fcs2"

// implement FlowData response structure

type FlowData struct {
	ID       int               `json:"id"` // internal ID
	Events   []float64         `json:"events"`
	MetaData *fcs2.Metadata `json:"metadata"`
}
