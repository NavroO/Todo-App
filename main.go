package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

type Note struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}

var notes []*Note

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Note
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		if err == io.EOF {
			http.Error(w, "Empty body", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	newNote := &Note{
		ID:     rand.Intn(1000000),
		Name:   data.Name,
		IsDone: data.IsDone,
	}

	notes = append(notes, newNote)
	json.NewEncoder(w).Encode(newNote)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Note
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		if err == io.EOF {
			http.Error(w, "Empty body", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	for i, note := range notes {
		if note.ID == data.ID {
			notes = append(notes[:i], notes[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Note
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		if err == io.EOF {
			http.Error(w, "Empty body", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	for _, note := range notes {
		if note.ID == data.ID {
			note.Name = data.Name
			note.IsDone = data.IsDone
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func main() {
	http.HandleFunc("/", getNote)
	http.HandleFunc("/addNote", addNote)
	http.HandleFunc("/deleteNote", deleteNote)
	http.HandleFunc("/updateNote", updateNote)
	http.ListenAndServe(":8080", nil)
}
