package handler

import (
	"encoding/json"
	"net/http"
	"notes-api/model"
	"notes-api/repository"
	"strconv"
	"strings"
)

type NoteHandler struct {
	repo repository.NoteRepository
}

func NewNoteHandler(repo repository.NoteRepository) http.Handler {
	return &NoteHandler{repo: repo}
}

func (h *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		path := r.URL.Path
		if strings.HasPrefix(path, "/note/") {
			h.GetNote(w, r)
		} else if path == "/notes" {
			h.GetNotes(w, r)
		} else {
			http.Error(w, "Not found", http.StatusNotFound)
		}
	case http.MethodPost:
		h.CreateNote(w, r)
	case http.MethodPut:
		h.UpdateNote(w, r)
	case http.MethodDelete:
		h.DeleteNote(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.repo.LoadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	noteId, _ := strconv.Atoi(r.PathValue("id"))
	note, err := h.repo.Load(noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	note, err := h.repo.Save(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteId, _ := strconv.Atoi(r.PathValue("id"))
	if err := h.repo.Delete(noteId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
