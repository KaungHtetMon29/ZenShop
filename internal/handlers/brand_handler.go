package handlers

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/internal/models"
	"net/http"
)

// BrandHandler handles brand-related requests
type BrandHandler struct {
	*Handler
}

// NewBrandHandler creates a new brand handler
func NewBrandHandler(h *Handler) *BrandHandler {
	return &BrandHandler{Handler: h}
}

// GetBrands returns all brands
func (h *BrandHandler) GetBrands(w http.ResponseWriter, r *http.Request) {
	var brands []models.Brand
	result := h.DB.Find(&brands)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve brands", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   brands,
		"status": "success",
		"count":  len(brands),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// CreateBrand creates a new brand
func (h *BrandHandler) CreateBrand(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Print(brand.Name)
	brand.CreatedAt = "2023-10-01"
	brand.UpdatedAt = "2023-10-01"
	result := h.DB.Create(&brand)
	fmt.Println(result)
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Hello, World!",
		"status":  "success",
		"data": map[string]interface{}{
			"id":   1,
			"name": "Sample User",
		},
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// UpdateBrand updates a brand by ID
func (h *BrandHandler) UpdateBrand(w http.ResponseWriter, r *http.Request) {
	// Extract brand ID from the URL path
	brandID := r.URL.Path[len("/brands/"):]
	if brandID == "" {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	var brand models.Brand
	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Attempt to update the brand by ID
	result := h.DB.Model(&models.Brand{}).Where("id = ?", brandID).Updates(brand)
	if result.Error != nil {
		http.Error(w, "Failed to update brand", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Brand not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Brand updated successfully",
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

// DeleteBrand deletes a brand by ID
func (h *BrandHandler) DeleteBrand(w http.ResponseWriter, r *http.Request) {
	// Extract brand ID from the URL path
	brandID := r.URL.Path[len("/brands/"):]
	if brandID == "" {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	// Attempt to delete the brand by ID
	result := h.DB.Delete(&models.Brand{}, brandID)
	if result.Error != nil {
		http.Error(w, "Failed to delete brand", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Brand not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Brand deleted successfully",
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
