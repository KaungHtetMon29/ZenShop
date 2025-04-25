package handlers

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/internal/models"
	"go_boilerplate/pkg"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// ProductHandler handles product-related requests
type ProductHandler struct {
	*Handler
}

// NewProductHandler creates a new product handler
func NewProductHandler(h *Handler) *ProductHandler {
	return &ProductHandler{Handler: h}
}

// GetProducts returns all products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	// Preload related brand and category data for each product
	result := h.DB.Preload("Brand").Preload("Category").Find(&products)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   products,
		"status": "success",
		"count":  len(products),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// GetFilteredProducts returns filtered products based on query parameters
func (h *ProductHandler) GetFilteredProducts(w http.ResponseWriter, r *http.Request) {
	// Parse URL query parameters
	queryParams := r.URL.Query()

	brandName := queryParams.Get("brand")
	categoryName := queryParams.Get("category")
	minPriceStr := queryParams.Get("min_price")
	maxPriceStr := queryParams.Get("max_price")

	// Initialize the database query
	query := h.DB.Model(&models.Product{}).
		Preload("Brand").
		Preload("Category")

	// Apply brand filter if provided
	if brandName != "" {
		// Join with brands table and filter by brand name
		query = query.Joins("JOIN brands ON brands.id = products.brand_id").
			Where("brands.name LIKE ?", "%"+brandName+"%")
	}

	// Apply category filter if provided
	if categoryName != "" {
		// Join with categories table and filter by category name
		query = query.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name LIKE ?", "%"+categoryName+"%")
	}

	// Apply price range filter if provided
	if minPriceStr != "" {
		minPrice, err := strconv.Atoi(minPriceStr)
		if err == nil {
			query = query.Where("products.price >= ?", minPrice)
		}
	}

	if maxPriceStr != "" {
		maxPrice, err := strconv.Atoi(maxPriceStr)
		if err == nil {
			query = query.Where("products.price <= ?", maxPrice)
		}
	}

	// Execute the query
	var products []models.Product
	result := query.Find(&products)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve filtered products: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return the filtered products as JSON response
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   products,
		"status": "success",
		"count":  len(products),
		"filters": map[string]interface{}{
			"brand":     brandName,
			"category":  categoryName,
			"min_price": minPriceStr,
			"max_price": maxPriceStr,
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

// CreateProduct creates a new product with image upload
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to parse form file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filebyte, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}
	fmt.Printf("File bytes: %d\n", len(filebyte))
	s3 := pkg.NewS3Config()
	url, err := s3.S3ImageUpload(filebyte, handler.Filename)
	if err != nil {
		http.Error(w, "Unable to upload image to S3", http.StatusInternalServerError)
		return
	}
	fmt.Printf("File URL: %s\n", url)

	brandID := r.FormValue("brandId")
	categoryID := r.FormValue("categoryId")
	name := r.FormValue("name")
	price := r.FormValue("price")
	stock := r.FormValue("stock")
	updateBy := r.FormValue("updateBy")

	// First, check if the brand_id exists and get the brand info
	var brandName string
	if brandID != "" {
		var brand models.Brand
		if err := h.DB.First(&brand, brandID).Error; err != nil {
			http.Error(w, "Invalid brand ID", http.StatusBadRequest)
			return
		}
		brandName = brand.Name
		fmt.Printf("Brand Name: %s\n", brandName)
	} else {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	// Create a map for inserting the product with all required fields
	productMap := map[string]interface{}{
		"brand_id":    brandID,
		"name":        name,
		"price":       price,
		"stock":       stock,
		"category_id": categoryID,
		"created_at":  "2023-10-01",
		"updated_at":  "2023-10-01",
		"update_by":   updateBy,
		"image_url":   url,
	}

	// Create the product using the map to ensure all fields are set
	result := h.DB.Model(&models.Product{}).Create(productMap)
	if result.Error != nil {
		http.Error(w, "Failed to create product: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Get the created product with its ID
	var createdProduct models.Product
	h.DB.First(&createdProduct, "name = ? AND brand_id = ?", name, brandID)

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product created successfully",
		"status":  "success",
		"data":    createdProduct,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// UpdateProduct updates a product by ID
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	productID := r.URL.Path[len("/products/"):]
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to parse form file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filebyte, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}
	fmt.Printf("File bytes: %d\n", len(filebyte))
	var product models.Product
	h.DB.First(&product, productID)
	fmt.Printf("Product Name:%s\n ", product.Name)
	var imageURL string
	if len(filebyte) != 0 {
		s3 := pkg.NewS3Config()
		err = s3.S3ImageDelete(strings.Split(product.ImageURL, "/")[len(strings.Split(product.ImageURL, "/"))-1])
		if err != nil {
			http.Error(w, "Unable to delete image on S3", http.StatusInternalServerError)
			return
		}
		url, err := s3.S3ImageUpload(filebyte, handler.Filename)
		imageURL = url
		if err != nil {
			http.Error(w, "Unable to upload image to S3", http.StatusInternalServerError)
			return
		}
		fmt.Printf("File URL: %s\n", url)
	} else {
		imageURL = product.ImageURL
	}
	brandID := r.FormValue("brandId")
	categoryID := r.FormValue("categoryId")
	name := r.FormValue("name")
	price := r.FormValue("price")
	stock := r.FormValue("stock")
	updateBy := r.FormValue("updateBy")
	updatedAt := "2023-10-01"

	result := h.DB.Model(&models.Product{}).Where("id = ?", productID).Updates(map[string]interface{}{
		"brand_id":    brandID,
		"category_id": categoryID,
		"name":        name,
		"price":       price,
		"stock":       stock,
		"update_by":   updateBy,
		"image_url":   imageURL,
		"updated_at":  updatedAt,
	})
	if result.Error != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product updated successfully",
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

// DeleteProduct deletes a product by ID
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/products/"):]
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	var product models.Product
	h.DB.First(&product, productID)
	fmt.Printf("Product Name:%s\n ", product.Name)
	s3 := pkg.NewS3Config()
	err := s3.S3ImageDelete(strings.Split(product.ImageURL, "/")[len(strings.Split(product.ImageURL, "/"))-1])
	if err != nil {
		http.Error(w, "Unable to delete image on S3", http.StatusInternalServerError)
		return
	}
	result := h.DB.Delete(&models.Product{}, productID)
	if result.Error != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product deleted successfully",
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
