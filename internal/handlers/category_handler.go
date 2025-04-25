package handlers

import (
	"encoding/json"
	"go_boilerplate/internal/models"
	"net/http"
)

// CategoryHandler handles category-related requests
type CategoryHandler struct {
	*Handler
}

// NewCategoryHandler creates a new category handler
func NewCategoryHandler(h *Handler) *CategoryHandler {
	return &CategoryHandler{Handler: h}
}

// GetCategories returns all categories
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	// Preload products associated with each category
	result := h.DB.Preload("Products").Find(&categories)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   categories,
		"status": "success",
		"count":  len(categories),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// CreateCategory creates a new category
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.CreatedAt = "2023-10-01"
	result := h.DB.Create(&category)
	if result.Error != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Category created successfully",
		"status":  "success",
		"data":    category,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// UpdateCategory updates a category by ID
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from the URL path
	categoryID := r.URL.Path[len("/categories/"):]
	if categoryID == "" {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Attempt to update the category by ID
	result := h.DB.Model(&models.Category{}).Where("id = ?", categoryID).Updates(category)
	if result.Error != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Category updated successfully",
		"status":  "success",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// DeleteCategory deletes a category by ID
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from the URL path
	categoryID := r.URL.Path[len("/categories/"):]
	if categoryID == "" {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	// Attempt to delete the category by ID
	result := h.DB.Delete(&models.Category{}, categoryID)
	if result.Error != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Category deleted successfully",
		"status":  "success",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
