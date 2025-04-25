package handlers

import (
	"encoding/json"
	"errors"
	"go_boilerplate/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ShippingHandler provides methods for dealing with shipping endpoints
type ShippingHandler struct {
	*Handler
}

// NewShippingHandler creates a new shipping handler
func NewShippingHandler(h *Handler) *ShippingHandler {
	return &ShippingHandler{
		Handler: h,
	}
}

// GetShippings handles GET requests to fetch all shipping records
func (h *ShippingHandler) GetShippings(w http.ResponseWriter, r *http.Request) {
	var shippings []models.Shipping

	result := h.DB.Find(&shippings)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve shipping records: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    shippings,
	})
}

// GetShipping handles GET requests to fetch a single shipping record by ID
func (h *ShippingHandler) GetShipping(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid shipping ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 32)
	if err != nil {
		http.Error(w, "Invalid shipping ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	var shipping models.Shipping
	result := h.DB.First(&shipping, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Shipping not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve shipping: "+result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    shipping,
	})
}

// CreateShipping handles POST requests to create a new shipping record
func (h *ShippingHandler) CreateShipping(w http.ResponseWriter, r *http.Request) {
	var shipping models.Shipping
	err := json.NewDecoder(r.Body).Decode(&shipping)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	shipping.CreatedAt = time.Now().Format(time.RFC3339)

	result := h.DB.Create(&shipping)
	if result.Error != nil {
		http.Error(w, "Failed to create shipping: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    shipping,
	})
}

// UpdateShipping handles PUT requests to update an existing shipping record
func (h *ShippingHandler) UpdateShipping(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid shipping ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 32)
	if err != nil {
		http.Error(w, "Invalid shipping ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	var shipping models.Shipping
	result := h.DB.First(&shipping, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Shipping not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve shipping: "+result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updatedShipping models.Shipping
	err = json.NewDecoder(r.Body).Decode(&updatedShipping)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields
	shipping.Address = updatedShipping.Address
	shipping.FirstName = updatedShipping.FirstName
	shipping.LastName = updatedShipping.LastName
	shipping.City = updatedShipping.City
	shipping.State = updatedShipping.State
	shipping.ZipCode = updatedShipping.ZipCode
	shipping.Country = updatedShipping.Country
	shipping.Email = updatedShipping.Email
	shipping.Phone = updatedShipping.Phone
	// Note: Not updating OrderID as this is a relationship field

	result = h.DB.Save(&shipping)
	if result.Error != nil {
		http.Error(w, "Failed to update shipping: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    shipping,
	})
}

// DeleteShipping handles DELETE requests to remove a shipping record
func (h *ShippingHandler) DeleteShipping(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid shipping ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(pathParts[len(pathParts)-1], 10, 32)
	if err != nil {
		http.Error(w, "Invalid shipping ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	var shipping models.Shipping
	result := h.DB.First(&shipping, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Shipping not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve shipping: "+result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	result = h.DB.Delete(&shipping)
	if result.Error != nil {
		http.Error(w, "Failed to delete shipping: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Shipping deleted successfully",
	})
}
