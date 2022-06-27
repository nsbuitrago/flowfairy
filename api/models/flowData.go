package models

import "github.com/nsbuitrago/fcs2/fcs2"

// implement FlowData response structure

type FlowData struct {
	ID       int               `json:"id"` // internal ID
	Events   []float64         `json:"events"`
	MetaData *fcs2.Metadata `json:"metadata"`
}

// implement map to the svelte ui structure 

type FlowData_Map struct {
	ID       int               `json:"id"` // internal ID
	Data   map[string]FlowData `json:"data"`
	MetaData *fcs2.Metadata    `json:"metadata"`
}
