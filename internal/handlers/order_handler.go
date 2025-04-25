package handlers

import (
	"encoding/json"
	"go_boilerplate/internal/models"
	"net/http"
)

// OrderHandler handles order-related requests
type OrderHandler struct {
	*Handler
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(h *Handler) *OrderHandler {
	return &OrderHandler{Handler: h}
}

// GetOrders returns all orders with their related data
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order

	// Preload related Payment and Shipping data for each order
	result := h.DB.Preload("Payment").Preload("Shipping").Find(&orders)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	// For each order, let's also load product details
	for i := range orders {
		h.DB.Preload("Product").Find(&orders[i].ProductPerOrder, "order_id = ?", orders[i].ID)
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   orders,
		"status": "success",
		"count":  len(orders),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// CreateOrder creates a new order with nested payment and shipping data
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Define a struct for the request format that matches frontend's data structure
	type OrderRequest struct {
		UserId    string `json:"UserId"`
		CreatedAt string `json:"CreatedAt"`
		Payment   struct {
			Amount          int    `json:"Amount"`
			Type            string `json:"Type"`
			CardholderName  string `json:"CardholderName"`
			CardNumberLast4 string `json:"CardNumberLast4"`
			ExpiryDate      string `json:"ExpiryDate"`
			CreatedAt       string `json:"CreatedAt"`
		} `json:"Payment"`
		Shipping struct {
			Address   string `json:"Address"`
			FirstName string `json:"FirstName"`
			LastName  string `json:"LastName"`
			City      string `json:"City"`
			State     string `json:"State"`
			ZipCode   string `json:"ZipCode"`
			Country   string `json:"Country"`
			Email     string `json:"Email"`
			Phone     string `json:"Phone"`
			CreatedAt string `json:"CreatedAt"`
		} `json:"Shipping"`
	}

	var orderReq OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Start a transaction to ensure all operations succeed or fail together
	tx := h.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// Create the order first
	order := models.Order{
		UserId:    orderReq.UserId,
		CreatedAt: orderReq.CreatedAt,
	}

	if result := tx.Create(&order); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create order: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Create payment record
	payment := models.Payment{
		OrderID:         order.ID,
		Amount:          orderReq.Payment.Amount,
		Type:            orderReq.Payment.Type,
		CardholderName:  orderReq.Payment.CardholderName,
		CardNumberLast4: orderReq.Payment.CardNumberLast4,
		ExpiryDate:      orderReq.Payment.ExpiryDate,
		CreatedAt:       orderReq.Payment.CreatedAt,
	}

	if result := tx.Create(&payment); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create payment record: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Create shipping record
	shipping := models.Shipping{
		OrderID:   order.ID,
		Address:   orderReq.Shipping.Address,
		FirstName: orderReq.Shipping.FirstName,
		LastName:  orderReq.Shipping.LastName,
		City:      orderReq.Shipping.City,
		State:     orderReq.Shipping.State,
		ZipCode:   orderReq.Shipping.ZipCode,
		Country:   orderReq.Shipping.Country,
		Email:     orderReq.Shipping.Email,
		Phone:     orderReq.Shipping.Phone,
		CreatedAt: orderReq.Shipping.CreatedAt,
	}

	if result := tx.Create(&shipping); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create shipping record: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Load the complete order with nested objects for the response
	var completeOrder models.Order
	h.DB.Preload("Payment").Preload("Shipping").First(&completeOrder, order.ID)

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Order created successfully",
		"status":  "success",
		"data":    completeOrder,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// UpdateOrder updates an order by ID with nested payment and shipping data
func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Path[len("/orders/"):]
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	// Define a struct for the request format that matches frontend's data structure
	type OrderRequest struct {
		UserId  string `json:"UserId"`
		Payment struct {
			Amount          int    `json:"Amount"`
			Type            string `json:"Type"`
			CardholderName  string `json:"CardholderName"`
			CardNumberLast4 string `json:"CardNumberLast4"`
			ExpiryDate      string `json:"ExpiryDate"`
		} `json:"Payment"`
		Shipping struct {
			Address   string `json:"Address"`
			FirstName string `json:"FirstName"`
			LastName  string `json:"LastName"`
			City      string `json:"City"`
			State     string `json:"State"`
			ZipCode   string `json:"ZipCode"`
			Country   string `json:"Country"`
			Email     string `json:"Email"`
			Phone     string `json:"Phone"`
		} `json:"Shipping"`
	}

	var orderReq OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Start a transaction to ensure all operations succeed or fail together
	tx := h.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// First, check if the order exists
	var existingOrder models.Order
	if err := tx.First(&existingOrder, orderID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Update order
	if result := tx.Model(&existingOrder).Updates(models.Order{
		UserId: orderReq.UserId,
	}); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	// Update payment
	if result := tx.Model(&models.Payment{}).Where("order_id = ?", orderID).Updates(models.Payment{
		Amount:          orderReq.Payment.Amount,
		Type:            orderReq.Payment.Type,
		CardholderName:  orderReq.Payment.CardholderName,
		CardNumberLast4: orderReq.Payment.CardNumberLast4,
		ExpiryDate:      orderReq.Payment.ExpiryDate,
	}); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to update payment", http.StatusInternalServerError)
		return
	}

	// Update shipping
	if result := tx.Model(&models.Shipping{}).Where("order_id = ?", orderID).Updates(models.Shipping{
		Address:   orderReq.Shipping.Address,
		FirstName: orderReq.Shipping.FirstName,
		LastName:  orderReq.Shipping.LastName,
		City:      orderReq.Shipping.City,
		State:     orderReq.Shipping.State,
		ZipCode:   orderReq.Shipping.ZipCode,
		Country:   orderReq.Shipping.Country,
		Email:     orderReq.Shipping.Email,
		Phone:     orderReq.Shipping.Phone,
	}); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to update shipping", http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	// Get the updated order with all related data for the response
	var updatedOrder models.Order
	h.DB.Preload("Payment").Preload("Shipping").First(&updatedOrder, orderID)

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Order updated successfully",
		"status":  "success",
		"data":    updatedOrder,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// DeleteOrder deletes an order by ID
func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Path[len("/orders/"):]
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	// Start a transaction
	tx := h.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// First, delete all product_per_order records associated with this order
	if err := tx.Where("order_id = ?", orderID).Delete(&models.ProductPerOrder{}).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete associated product orders: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Then, delete the shipping record
	if err := tx.Where("order_id = ?", orderID).Delete(&models.Shipping{}).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete shipping record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the payment record
	if err := tx.Where("order_id = ?", orderID).Delete(&models.Payment{}).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete payment record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Finally, delete the order itself
	result := tx.Delete(&models.Order{}, orderID)
	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete order: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Order deleted successfully",
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

// GetProductPerOrders returns all product per orders
func (h *OrderHandler) GetProductPerOrders(w http.ResponseWriter, r *http.Request) {
	var productOrders []models.ProductPerOrder
	result := h.DB.Find(&productOrders)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve product orders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   productOrders,
		"status": "success",
		"count":  len(productOrders),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// CreateProductPerOrder creates a new product per order
func (h *OrderHandler) CreateProductPerOrder(w http.ResponseWriter, r *http.Request) {
	var productOrder models.ProductPerOrder
	if err := json.NewDecoder(r.Body).Decode(&productOrder); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	productOrder.CreatedAt = "2023-10-01"
	result := h.DB.Create(&productOrder)
	if result.Error != nil {
		http.Error(w, "Failed to create product order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product order created successfully",
		"status":  "success",
		"data":    productOrder,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// DeleteProductPerOrder deletes a product per order by ID
func (h *OrderHandler) DeleteProductPerOrder(w http.ResponseWriter, r *http.Request) {
	productOrderID := r.URL.Path[len("/product-orders/"):]
	if productOrderID == "" {
		http.Error(w, "Product Order ID is required", http.StatusBadRequest)
		return
	}

	result := h.DB.Delete(&models.ProductPerOrder{}, productOrderID)
	if result.Error != nil {
		http.Error(w, "Failed to delete product order", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Product order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product order deleted successfully",
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
