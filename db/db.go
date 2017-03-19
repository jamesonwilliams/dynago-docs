package db

import (
	"github.com/jamesonwilliams/dynago-docs/model"
)

type Database interface {
	RetrieveDocument(id int) model.Document
	RetrieveDocuments() []model.Document
	StoreDocument(d model.Document) model.Document
	DeleteDocument(id int) error
}
