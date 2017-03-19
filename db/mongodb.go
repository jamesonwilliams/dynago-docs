package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jamesonwilliams/dynago-docs/model"
)

const databaseHostname = "localhost"
const databaseName = "documents"
const documentsCollectionKey = "document"

// MongoDatabase Encapsulates a connection to a database.
type MongoDatabase struct {
	session *mgo.Session
}

func (mdb *MongoDatabase) StoreDocument(document model.Document) (model.Document, error) {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	if _, err := mdb.RetrieveDocument(document.Id); err == nil {
		return model.Document{}, err
	}
	collection := mdb.session.DB(databaseName).C(documentsCollectionKey)
	err := collection.Insert(collection)
	return document, err
}

func (mdb *MongoDatabase) RetrieveDocuments() ([]model.Document, error) {
	collection := mdb.session.DB(databaseName).C(documentsCollectionKey)
	var documents []model.Document
	err := collection.Find(nil).All(&documents)
	return documents, err
}

// RetrieveDocument get data from a document.
func (mdb *MongoDatabase) RetrieveDocument(id int) (result model.Document, err error) {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	documents := mdb.session.DB(databaseName).C(documentsCollectionKey)
	err = documents.Find(bson.M{"documentId": id}).One(&result)
	return result, err
}

func (mdb *MongoDatabase) DeleteDocument(id int) error {
	return nil
}

// GetSession return a new session if there is no previous one.
func (mdb *MongoDatabase) GetSession() *mgo.Session {
	if mdb.session != nil {
		return mdb.session.Copy()
	}
	session, err := mgo.Dial(databaseHostname)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
