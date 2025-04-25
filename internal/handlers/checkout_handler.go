package handlers

import (
	"encoding/json"
	"go_boilerplate/internal/models"
	"net/http"
	"time"
)

// CheckoutHandler handles checkout-related requests
type CheckoutHandler struct {
	*Handler
}

// NewCheckoutHandler creates a new checkout handler
func NewCheckoutHandler(h *Handler) *CheckoutHandler {
	return &CheckoutHandler{Handler: h}
}

// Checkout request structure that matches the frontend request
type CheckoutRequest struct {
	Shipping struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Address   string `json:"address"`
		City      string `json:"city"`
		State     string `json:"state"`
		ZipCode   string `json:"zipCode"`
		Country   string `json:"country"`
		Email     string `json:"email,omitempty"`
		Phone     string `json:"phone,omitempty"`
	} `json:"shipping"`
	Payment struct {
		CardholderName  string `json:"cardholderName"`
		CardNumberLast4 string `json:"cardNumberLast4"`
		ExpiryDate      string `json:"expiryDate"`
	} `json:"payment"`
	Order struct {
		Items []struct {
			ID       uint   `json:"id"`
			Name     string `json:"name"`
			Price    int    `json:"price"`
			Quantity int    `json:"quantity"`
		} `json:"items"`
		TotalItems  int `json:"totalItems"`
		Subtotal    int `json:"subtotal"`
		ShippingFee int `json:"shippingFee"`
		Total       int `json:"total"`
	} `json:"order"`
}

// ProcessCheckout handles the checkout process
func (h *CheckoutHandler) ProcessCheckout(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var checkoutReq CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&checkoutReq); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Start a transaction
	tx := h.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// Create order record
	currentDate := time.Now().Format("2006-01-02")
	order := models.Order{
		UserId:    "guest", // You may want to update this with actual user ID if authenticated
		CreatedAt: currentDate,
	}

	if result := tx.Create(&order); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create order: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Create shipping record
	shipping := models.Shipping{
		OrderID:   order.ID,
		FirstName: checkoutReq.Shipping.FirstName,
		LastName:  checkoutReq.Shipping.LastName,
		Address:   checkoutReq.Shipping.Address,
		City:      checkoutReq.Shipping.City,
		State:     checkoutReq.Shipping.State,
		ZipCode:   checkoutReq.Shipping.ZipCode,
		Country:   checkoutReq.Shipping.Country,
		Email:     checkoutReq.Shipping.Email,
		Phone:     checkoutReq.Shipping.Phone,
		CreatedAt: currentDate,
	}

	if result := tx.Create(&shipping); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create shipping record: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Create payment record
	payment := models.Payment{
		OrderID:         order.ID,
		Amount:          checkoutReq.Order.Total,
		CardholderName:  checkoutReq.Payment.CardholderName,
		CardNumberLast4: checkoutReq.Payment.CardNumberLast4,
		ExpiryDate:      checkoutReq.Payment.ExpiryDate,
		CreatedAt:       currentDate,
		Type:            "Credit Card", // Assuming credit card payment
	}

	if result := tx.Create(&payment); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create payment record: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Create product per order records for each item
	for _, item := range checkoutReq.Order.Items {
		productOrder := models.ProductPerOrder{
			OrderID:   order.ID,
			ProductID: item.ID,
			Quantity:  item.Quantity,
			CreatedAt: currentDate,
		}

		if result := tx.Create(&productOrder); result.Error != nil {
			tx.Rollback()
			http.Error(w, "Failed to create product order record: "+result.Error.Error(), http.StatusInternalServerError)
			return
		}

		// Update product stock
		var product models.Product
		if result := tx.First(&product, item.ID); result.Error != nil {
			tx.Rollback()
			http.Error(w, "Failed to find product: "+result.Error.Error(), http.StatusInternalServerError)
			return
		}

		// Ensure we have enough stock
		if product.Stock < item.Quantity {
			tx.Rollback()
			http.Error(w, "Insufficient stock for product: "+product.Name, http.StatusBadRequest)
			return
		}

		// Update the stock
		product.Stock -= item.Quantity
		if result := tx.Save(&product); result.Error != nil {
			tx.Rollback()
			http.Error(w, "Failed to update product stock: "+result.Error.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch the complete order with all related data for the response
	var completeOrder models.Order
	h.DB.Preload("Payment").Preload("Shipping").Preload("ProductPerOrder.Product").First(&completeOrder, order.ID)

	// Return success response
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Checkout completed successfully",
		"status":  "success",
		"data":    completeOrder,
		"orderID": order.ID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
