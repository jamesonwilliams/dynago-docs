package db

import (
	"fmt"
	"github.com/jamesonwilliams/dynago-docs/model"
)

type InMemDatabase struct {
	currentId int
	documents []model.Document
}

func (db *InMemDatabase) RetrieveDocuments() ([]model.Document, error) {
	return db.documents, nil
}

func (db *InMemDatabase) RetrieveDocument(id string) (model.Document, error) {
	for _, t := range db.documents {
		if t.Id == id {
			return t, nil
		}
	}
	// return empty Document if not found
	return model.Document{}, nil
}

//this is bad, I don't think it passes race condtions
func (db *InMemDatabase) StoreDocument(t model.Document) (model.Document, error) {
	db.currentId += 1
	t.Id = string(db.currentId)
	db.documents = append(db.documents, t)
	return t, nil
}

func (db *InMemDatabase) DeleteDocument(id string) error {
	for i, t := range db.documents {
		if t.Id == id {
			db.documents = append(db.documents[:i], db.documents[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Document with id of %d to delete", id)
}
