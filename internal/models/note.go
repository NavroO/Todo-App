package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
	UserID int
	User   User
}
