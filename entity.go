package main

import (
	"github.com/google/uuid"
)

// Entity Struct
type Student struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	GPA         float64   `json:"gpa"`
	IsGraduate  bool      `json:"is_graduate"`
}