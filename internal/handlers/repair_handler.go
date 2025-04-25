package handlers

import (
	"encoding/json"
	"go_boilerplate/internal/models"
	"net/http"
)

// RepairHandler handles repair-related requests
type RepairHandler struct {
	*Handler
}

// NewRepairHandler creates a new repair handler
func NewRepairHandler(h *Handler) *RepairHandler {
	return &RepairHandler{Handler: h}
}

// GetRepairs returns all repairs
func (h *RepairHandler) GetRepairs(w http.ResponseWriter, r *http.Request) {
	var repairs []models.Repair
	result := h.DB.Find(&repairs)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve repairs", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   repairs,
		"status": "success",
		"count":  len(repairs),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// CreateRepair creates a new repair
func (h *RepairHandler) CreateRepair(w http.ResponseWriter, r *http.Request) {
	var repair models.Repair
	if err := json.NewDecoder(r.Body).Decode(&repair); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repair.CreatedAt = "2023-10-01"
	repair.UpdatedAt = "2023-10-01"
	result := h.DB.Create(&repair)
	if result.Error != nil {
		http.Error(w, "Failed to create repair", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Repair created successfully",
		"status":  "success",
		"data":    repair,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// UpdateRepair updates a repair by ID
func (h *RepairHandler) UpdateRepair(w http.ResponseWriter, r *http.Request) {
	repairID := r.URL.Path[len("/repairs/"):]
	if repairID == "" {
		http.Error(w, "Repair ID is required", http.StatusBadRequest)
		return
	}

	var repair models.Repair
	if err := json.NewDecoder(r.Body).Decode(&repair); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repair.UpdatedAt = "2023-10-01"
	result := h.DB.Model(&models.Repair{}).Where("id = ?", repairID).Updates(repair)
	if result.Error != nil {
		http.Error(w, "Failed to update repair", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Repair not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Repair updated successfully",
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

// DeleteRepair deletes a repair by ID
func (h *RepairHandler) DeleteRepair(w http.ResponseWriter, r *http.Request) {
	repairID := r.URL.Path[len("/repairs/"):]
	if repairID == "" {
		http.Error(w, "Repair ID is required", http.StatusBadRequest)
		return
	}

	result := h.DB.Delete(&models.Repair{}, repairID)
	if result.Error != nil {
		http.Error(w, "Failed to delete repair", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Repair not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Repair deleted successfully",
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

// GetRepairStatuses returns all repair statuses
func (h *RepairHandler) GetRepairStatuses(w http.ResponseWriter, r *http.Request) {
	var repairStatuses []models.RepairStatus
	result := h.DB.Find(&repairStatuses)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve repair statuses", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   repairStatuses,
		"status": "success",
		"count":  len(repairStatuses),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// CreateRepairStatus creates a new repair status
func (h *RepairHandler) CreateRepairStatus(w http.ResponseWriter, r *http.Request) {
	var repairStatus models.RepairStatus
	if err := json.NewDecoder(r.Body).Decode(&repairStatus); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repairStatus.UpdatedAt = "2023-10-01"
	result := h.DB.Create(&repairStatus)
	if result.Error != nil {
		http.Error(w, "Failed to create repair status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Repair status created successfully",
		"status":  "success",
		"data":    repairStatus,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// UpdateRepairStatus updates a repair status by ID
func (h *RepairHandler) UpdateRepairStatus(w http.ResponseWriter, r *http.Request) {
	statusID := r.URL.Path[len("/repair-statuses/"):]
	if statusID == "" {
		http.Error(w, "Repair Status ID is required", http.StatusBadRequest)
		return
	}

	var repairStatus models.RepairStatus
	if err := json.NewDecoder(r.Body).Decode(&repairStatus); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repairStatus.UpdatedAt = "2023-10-01"
	result := h.DB.Model(&models.RepairStatus{}).Where("id = ?", statusID).Updates(repairStatus)
	if result.Error != nil {
		http.Error(w, "Failed to update repair status", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Repair status not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Repair status updated successfully",
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

// DeleteRepairStatus deletes a repair status by ID
func (h *RepairHandler) DeleteRepairStatus(w http.ResponseWriter, r *http.Request) {
	statusID := r.URL.Path[len("/repair-statuses/"):]
	if statusID == "" {
		http.Error(w, "Repair Status ID is required", http.StatusBadRequest)
		return
	}

	result := h.DB.Delete(&models.RepairStatus{}, statusID)
	if result.Error != nil {
		http.Error(w, "Failed to delete repair status", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Repair status not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Repair status deleted successfully",
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
