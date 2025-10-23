package main

import (
	"fmt"
	"net/http"
	"notes-api/handler"
	"notes-api/repository"
)

func main() {

	repo := repository.GetNoteRepository("sqlite3")
	noteHandler := handler.NewNoteHandler(repo)
	server := NewRouter(noteHandler)

	fmt.Println("Listening on https://localhost:8080")
	http.ListenAndServe(":8080", server)
}
