package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"eventbus/internal/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var notes []models.Note
	if err := db.Find(&notes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notes)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Note
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		if err == io.EOF {
			http.Error(w, "Empty body", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	if err := db.Create(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "noteID")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&models.Note{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data models.Note
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		if err == io.EOF {
			http.Error(w, "Empty body", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	var note models.Note
	if err := db.First(&note, data.ID).Error; err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	note.Name = data.Name
	note.IsDone = data.IsDone

	if err := db.Save(&note).Error; err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Note{}, &models.User{})

	r := chi.NewRouter()

	r.Route("/notes", func(r chi.Router) {
		r.Post("/", addNote)
		r.Get("/", getNotes)
		r.Route("/{noteID}", func(r chi.Router) {
			r.Delete("/", deleteNote)
			r.Patch("/", updateNote)
		})
	})

	http.ListenAndServe(":8080", r)
}
