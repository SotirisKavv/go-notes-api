package repository

import (
	"database/sql"
	"fmt"
	"notes-api/database"
	"notes-api/model"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	db, err := database.InitSQLite("database/notes.db")
	if err != nil {
		fmt.Printf("Error initializing DB: %v\n", err)
		return nil
	}
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Load(id int) (model.Note, error) {
	note := model.Note{}
	err := r.db.QueryRow(`SELECT id, title, body, user_id FROM notes WHERE id = ?`, id).Scan(&note.ID, &note.Title, &note.Body, &note.UserID)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *SQLiteRepository) LoadAll() ([]model.Note, error) {
	var notes []model.Note
	rows, err := r.db.Query(`SELECT id, title, body, user_id FROM notes`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := model.Note{}
		rows.Scan(&note.ID, &note.Title, &note.Body, &note.UserID)
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *SQLiteRepository) Save(note model.Note) (model.Note, error) {
	res, err := r.db.Exec("INSERT INTO notes(title, body, user_id) VALUES(?, ?, ?)", note.Title, note.Body, note.UserID)
	if err != nil {
		return model.Note{}, err
	}
	id, _ := res.LastInsertId()
	note.ID = int(id)
	return note, nil
}

func (r *SQLiteRepository) SaveAll(notes []model.Note) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO notes(title, body, user_id) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, note := range notes {
		_, err := stmt.Exec(note.Title, note.Body, note.UserID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SQLiteRepository) Update(note model.Note) error {
	_, err := r.db.Exec(`UPDATE notes SET title=?, body=?, user_id=? WHERE id=?`, note.Title, note.Body, note.UserID, note.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
