package handlers

import (
	"gorm.io/gorm"
)

// Handler is the base handler that contains shared DB connection
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new handler with DB connection
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		DB: db,
	}
}
