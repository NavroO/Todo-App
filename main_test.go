package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Note{})
	db.Create(&Note{Name: "Test Note 1", IsDone: false})
	db.Create(&Note{Name: "Test Note 2", IsDone: true})
	return db
}

func TestGetNote(t *testing.T) {
	db = setupTestDB()

	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	w := httptest.NewRecorder()

	getNote(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", resp.StatusCode)
	}

	var got []Note
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("cannot decode response: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 notes, got %d", len(got))
	}

	if got[0].Name != "Test Note 1" || got[1].Name != "Test Note 2" {
		t.Errorf("unexpected note names: %+v", got)
	}
}

func TestAddNote(t *testing.T) {
	db = setupTestDB()

	note := Note{
		Name:   "New testing note",
		IsDone: false,
	}
	body, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/addNote", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	addNote(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", resp.StatusCode)
	}
}
