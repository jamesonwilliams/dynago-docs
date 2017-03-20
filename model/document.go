package model

type Document struct {
	Id        string `json:"documentId"`
	Published string `json:"published"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
}
