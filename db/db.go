package db

import (
	"github.com/jamesonwilliams/dynago-docs/model"
)

type Database interface {
	RetrieveDocument(id string) (model.Document, error)
	RetrieveDocuments() ([]model.Document, error)
	StoreDocument(d model.Document) (model.Document, error)
	DeleteDocument(id string) error
}
