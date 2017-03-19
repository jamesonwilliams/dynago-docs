package db

import (
	"fmt"
	"github.com/jamesonwilliams/dynago-docs/model"
)

type InMemDatabase struct {
	currentId int
	documents []model.Document
}

func (db *InMemDatabase) RetrieveDocuments() []model.Document {
	return db.documents
}

func (db *InMemDatabase) RetrieveDocument(id int) model.Document {
	for _, t := range db.documents {
		if t.Id == id {
			return t
		}
	}
	// return empty Document if not found
	return model.Document{}
}

//this is bad, I don't think it passes race condtions
func (db *InMemDatabase) StoreDocument(t model.Document) model.Document {
	db.currentId += 1
	t.Id = db.currentId
	db.documents = append(db.documents, t)
	return t
}

func (db *InMemDatabase) DeleteDocument(id int) error {
	for i, t := range db.documents {
		if t.Id == id {
			db.documents = append(db.documents[:i], db.documents[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Document with id of %d to delete", id)
}
