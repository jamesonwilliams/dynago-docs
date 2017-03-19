package main

import "fmt"

var currentId int

var documents Documents

// Give us some seed data
func init() {
	RepoCreateDocument(Document{Name: "Write presentation"})
	RepoCreateDocument(Document{Name: "Host meetup"})
}

func RepoFindDocument(id int) Document {
	for _, t := range documents {
		if t.Id == id {
			return t
		}
	}
	// return empty Document if not found
	return Document{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateDocument(t Document) Document {
	currentId += 1
	t.Id = currentId
	documents = append(documents, t)
	return t
}

func RepoDestroyDocument(id int) error {
	for i, t := range documents {
		if t.Id == id {
			documents = append(documents[:i], documents[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Document with id of %d to delete", id)
}
