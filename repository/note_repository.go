package repository

import (
	"notes-api/model"
)

type NoteRepository interface {
	Load(id int) (model.Note, error)
	LoadAll() ([]model.Note, error)
	Save(note model.Note) (model.Note, error)
	SaveAll(notes []model.Note) error
	Update(note model.Note) error
	Delete(id int) error
}

func GetNoteRepository(repoType string) NoteRepository {
	switch repoType {
	case "memory":
		return NewMemoryRepository()
	case "sqlite3":
		return NewSQLiteRepository()
	default:
		return nil
	}
}
