package model

import "time"

type Document struct {
	Id        string    `json:"documentId"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
