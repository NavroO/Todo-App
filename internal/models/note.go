package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Name   string `json:"name" validate:"required,min=2,max=50"`
	IsDone bool   `json:"is_done"`
	UserID int
	User   User
}
