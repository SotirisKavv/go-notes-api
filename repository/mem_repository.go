package repository

import (
	"fmt"
	"notes-api/model"
)

type MemoryRepository struct {
	notes  map[int]model.Note
	nextID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		notes:  make(map[int]model.Note),
		nextID: 0,
	}
}

func (r *MemoryRepository) Load(id int) (model.Note, error) {
	note, ok := r.notes[id]
	if !ok {
		return model.Note{}, fmt.Errorf("Note with id %d not found", id)
	}

	return note, nil
}

func (r *MemoryRepository) LoadAll() ([]model.Note, error) {
	var notes []model.Note

	for _, note := range r.notes {
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *MemoryRepository) Save(note model.Note) (model.Note, error) {
	note.ID = r.nextID
	r.notes[r.nextID] = note
	r.nextID++
	return note, nil
}

func (r *MemoryRepository) SaveAll(notes []model.Note) error {
	for _, note := range notes {
		r.notes[note.ID] = note
	}

	return nil
}

func (r *MemoryRepository) Update(note model.Note) error {
	id := note.ID
	_, ok := r.notes[id]
	if !ok {
		return fmt.Errorf("Note with id %d not found", id)
	}

	r.notes[id] = note

	return nil
}

func (r *MemoryRepository) Delete(id int) error {
	delete(r.notes, id)
	return nil
}
