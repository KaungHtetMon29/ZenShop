package routes

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/internal/middleware"
	"go_boilerplate/internal/models"
	"go_boilerplate/pkg"
	"io"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type router struct {
	routes map[string]map[string]http.HandlerFunc
	db     *gorm.DB
}

func NewRouter(db *gorm.DB) *router {
	return &router{
		routes: make(map[string]map[string]http.HandlerFunc),
		db:     db,
	}
}

func InitializeRoutes(db *gorm.DB) *router {
	r := NewRouter(db)

	// Define the final handler that will return the data
	// finalHandler := http.HandlerFunc(r.testHandler)
	addUserHandler := http.HandlerFunc(r.inputData)

	// Brand handlers
	addBrandHandler := http.HandlerFunc(r.inputBrand)
	getBrandHandler := http.HandlerFunc(r.getBrand)
	deleteBrandHandler := http.HandlerFunc(r.deleteBrand)
	updateBrandHandler := http.HandlerFunc(r.updateBrand)

	// Category handlers
	addCategoryHandler := http.HandlerFunc(r.inputCategory)
	getCategoryHandler := http.HandlerFunc(r.getCategory)
	deleteCategoryHandler := http.HandlerFunc(r.deleteCategory)
	updateCategoryHandler := http.HandlerFunc(r.updateCategory)

	// Product handlers
	getProductHandler := http.HandlerFunc(r.getProduct)
	addProductHandler := http.HandlerFunc(r.inputProduct)
	updateProductHandler := http.HandlerFunc(r.updateProduct)
	deleteProductHandler := http.HandlerFunc(r.deleteProduct)

	// Order handlers
	getOrderHandler := http.HandlerFunc(r.getOrder)
	addOrderHandler := http.HandlerFunc(r.inputOrder)
	updateOrderHandler := http.HandlerFunc(r.updateOrder)
	deleteOrderHandler := http.HandlerFunc(r.deleteOrder)

	// Repair handlers
	getRepairHandler := http.HandlerFunc(r.getRepair)
	addRepairHandler := http.HandlerFunc(r.inputRepair)
	updateRepairHandler := http.HandlerFunc(r.updateRepair)
	deleteRepairHandler := http.HandlerFunc(r.deleteRepair)

	// RepairStatus handlers
	getRepairStatusHandler := http.HandlerFunc(r.getRepairStatus)
	addRepairStatusHandler := http.HandlerFunc(r.inputRepairStatus)
	updateRepairStatusHandler := http.HandlerFunc(r.updateRepairStatus)
	deleteRepairStatusHandler := http.HandlerFunc(r.deleteRepairStatus)

	// ProductUpdateHistory handlers
	getProductUpdateHistoryHandler := http.HandlerFunc(r.getProductUpdateHistory)
	addProductUpdateHistoryHandler := http.HandlerFunc(r.inputProductUpdateHistory)
	deleteProductUpdateHistoryHandler := http.HandlerFunc(r.deleteProductUpdateHistory)

	// Payment handlers
	getPaymentHandler := http.HandlerFunc(r.getPayment)
	addPaymentHandler := http.HandlerFunc(r.inputPayment)
	updatePaymentHandler := http.HandlerFunc(r.updatePayment)
	deletePaymentHandler := http.HandlerFunc(r.deletePayment)

	// Shipping handlers
	getShippingHandler := http.HandlerFunc(r.getShipping)
	addShippingHandler := http.HandlerFunc(r.inputShipping)
	updateShippingHandler := http.HandlerFunc(r.updateShipping)
	deleteShippingHandler := http.HandlerFunc(r.deleteShipping)

	// ProductPerOrder handlers
	getProductPerOrderHandler := http.HandlerFunc(r.getProductPerOrder)
	addProductPerOrderHandler := http.HandlerFunc(r.inputProductPerOrder)
	deleteProductPerOrderHandler := http.HandlerFunc(r.deleteProductPerOrder)

	// Apply middleware to the final handler
	r.AddRoute("GET", "/users", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(addUserHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/users", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(addUserHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("GET", "/brands", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(getBrandHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/brands", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(addBrandHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/brands/:id", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(deleteBrandHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/brands/:id", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(updateBrandHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Category routes
	r.AddRoute("GET", "/categories", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(getCategoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/categories", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(addCategoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/categories/:id", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(deleteCategoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/categories/:id", func(w http.ResponseWriter, req *http.Request) {
		// Create middleware chain and execute it
		handler := middleware.SetHandler(updateCategoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Product routes
	r.AddRoute("GET", "/products", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getProductHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/products", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addProductHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("PUT", "/products/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(updateProductHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/products/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteProductHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Order routes
	r.AddRoute("GET", "/orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("PUT", "/orders/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(updateOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/orders/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Repair routes
	r.AddRoute("GET", "/repairs", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getRepairHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/repairs", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addRepairHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("PUT", "/repairs/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(updateRepairHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/repairs/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteRepairHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// RepairStatus routes
	r.AddRoute("GET", "/repair-statuses", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getRepairStatusHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/repair-statuses", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addRepairStatusHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("PUT", "/repair-statuses/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(updateRepairStatusHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/repair-statuses/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteRepairStatusHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// ProductUpdateHistory routes
	r.AddRoute("GET", "/product-histories", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getProductUpdateHistoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/product-histories", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addProductUpdateHistoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/product-histories/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteProductUpdateHistoryHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Payment routes
	r.AddRoute("GET", "/payments", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getPaymentHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/payments", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addPaymentHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("PUT", "/payments/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(updatePaymentHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/payments/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deletePaymentHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Shipping routes
	r.AddRoute("GET", "/shippings", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getShippingHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/shippings", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addShippingHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("PUT", "/shippings/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(updateShippingHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/shippings/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteShippingHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// ProductPerOrder routes
	r.AddRoute("GET", "/product-orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(getProductPerOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("POST", "/product-orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(addProductPerOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})
	r.AddRoute("DELETE", "/product-orders/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(deleteProductPerOrderHandler)
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	return r
}

func (rt *router) updateBrand(w http.ResponseWriter, r *http.Request) {
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
	result := rt.db.Model(&models.Brand{}).Where("id = ?", brandID).Updates(brand)
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

func (rt *router) deleteBrand(w http.ResponseWriter, r *http.Request) {
	// Extract brand ID from the URL path
	brandID := r.URL.Path[len("/brands/"):]
	if brandID == "" {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	// Attempt to delete the brand by ID
	result := rt.db.Delete(&models.Brand{}, brandID)
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

func (rt *router) getBrand(w http.ResponseWriter, r *http.Request) {
	var brands []models.Brand
	result := rt.db.Find(&brands)
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

func (rt *router) inputBrand(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Print(brand.Name)
	brand.CreatedAt = "2023-10-01"
	brand.UpdatedAt = "2023-10-01"
	result := rt.db.Create(&brand)
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

func (rt *router) inputData(w http.ResponseWriter, r *http.Request) {
	brand := models.Brand{Name: "test2", CreatedAt: "2023-10-01", UpdatedAt: "2023-10-01"}
	result := rt.db.Create(&brand)
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

func testmw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("🔥 testmw() start")
		// Do middleware processing here

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
func (rt *router) testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🔥 testHandler() start")
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

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// First try exact match
	if handlers, ok := r.routes[path]; ok {
		if handler, ok := handlers[method]; ok {
			handler(w, req)
			return
		}
	}

	// If no exact match, try to match patterns with parameters
	for routePath, handlers := range r.routes {
		if handler, ok := handlers[method]; ok {
			// Check if the route contains a parameter (e.g., /:id)
			if pathMatches(routePath, path) {
				handler(w, req)
				return
			}
		}
	}

	http.NotFound(w, req)
}

// pathMatches checks if a URL path matches a route pattern with parameters
func pathMatches(pattern, path string) bool {
	// Split the pattern and path into segments
	patternParts := splitPath(pattern)
	pathParts := splitPath(path)

	// If they have different number of segments, they don't match
	if len(patternParts) != len(pathParts) {
		return false
	}

	// Check each segment
	for i, part := range patternParts {
		// If this segment is a parameter (starts with :), it matches anything
		if len(part) > 0 && part[0] == ':' {
			continue
		}

		// Otherwise, segments must match exactly
		if part != pathParts[i] {
			return false
		}
	}

	return true
}

// splitPath splits a URL path into segments
func splitPath(path string) []string {
	// Remove leading slash if present
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// Handle empty path
	if path == "" {
		return []string{}
	}

	return strings.Split(path, "/")
}

func (r *router) AddRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}

// Category CRUD handlers
func (rt *router) getCategory(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	result := rt.db.Find(&categories)
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

func (rt *router) inputCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.CreatedAt = "2023-10-01"
	result := rt.db.Create(&category)
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

func (rt *router) updateCategory(w http.ResponseWriter, r *http.Request) {
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
	result := rt.db.Model(&models.Category{}).Where("id = ?", categoryID).Updates(category)
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

func (rt *router) deleteCategory(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from the URL path
	categoryID := r.URL.Path[len("/categories/"):]
	if categoryID == "" {
		http.Error(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	// Attempt to delete the category by ID
	result := rt.db.Delete(&models.Category{}, categoryID)
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

// Product CRUD handlers
func (rt *router) getProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	// Preload related brand and category data for each product
	result := rt.db.Preload("Brand").Preload("Category").Find(&products)
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

func (rt *router) inputProduct(w http.ResponseWriter, r *http.Request) {
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

	// var product models.Product
	// if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }

	// First, check if the brand_id exists and get the brand info
	var brandName string
	if brandID != "" {
		var brand models.Brand
		if err := rt.db.First(&brand, brandID).Error; err != nil {
			http.Error(w, "Invalid brand ID", http.StatusBadRequest)
			return
		}
		brandName = brand.Name
		fmt.Printf("Brand Name: %s\n", brandName)
	} else {
		http.Error(w, "Brand ID is required", http.StatusBadRequest)
		return
	}

	// Check if category exists
	// if categoryID != "" {
	// 	var category models.Category
	// 	if err := rt.db.First(&category, categoryID).Error; err != nil {
	// 		http.Error(w, "Invalid category ID", http.StatusBadRequest)
	// 		return
	// 	}
	// }

	// Create a map for inserting the product with all required fields
	productMap := map[string]interface{}{
		"brand_id": brandID,
		// "brand":       brandName, // This is the key addition - setting the brand text field
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
	result := rt.db.Model(&models.Product{}).Create(productMap)
	if result.Error != nil {
		http.Error(w, "Failed to create product: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// // Get the created product with its ID
	var createdProduct models.Product
	rt.db.First(&createdProduct, "name = ? AND brand_id = ?", name, brandID)

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

func (rt *router) updateProduct(w http.ResponseWriter, r *http.Request) {
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
	// if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }
	rt.db.First(&product, productID)
	fmt.Printf("Product Name:%s\n " + product.Name)
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

	result := rt.db.Model(&models.Product{}).Where("id = ?", productID).Updates(map[string]interface{}{
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

	// if result.RowsAffected == 0 {
	// 	http.Error(w, "Product not found", http.StatusNotFound)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// response := map[string]interface{}{
	// 	"message": "Product updated successfully",
	// 	"status":  "success",
	// }
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonResponse)
}

func (rt *router) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/products/"):]
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	var product models.Product
	// if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }
	rt.db.First(&product, productID)
	fmt.Printf("Product Name:%s\n " + product.Name)
	s3 := pkg.NewS3Config()
	err := s3.S3ImageDelete(strings.Split(product.ImageURL, "/")[len(strings.Split(product.ImageURL, "/"))-1])
	if err != nil {
		http.Error(w, "Unable to delete image on S3", http.StatusInternalServerError)
		return
	}
	result := rt.db.Delete(&models.Product{}, productID)
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

// Order CRUD handlers
func (rt *router) getOrder(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	result := rt.db.Find(&orders)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
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

func (rt *router) inputOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	order.CreatedAt = "2023-10-01"
	result := rt.db.Create(&order)
	if result.Error != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Order created successfully",
		"status":  "success",
		"data":    order,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) updateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Path[len("/orders/"):]
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := rt.db.Model(&models.Order{}).Where("id = ?", orderID).Updates(order)
	if result.Error != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Order updated successfully",
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

func (rt *router) deleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Path[len("/orders/"):]
	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.Order{}, orderID)
	if result.Error != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
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

// Repair CRUD handlers
func (rt *router) getRepair(w http.ResponseWriter, r *http.Request) {
	var repairs []models.Repair
	result := rt.db.Find(&repairs)
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

func (rt *router) inputRepair(w http.ResponseWriter, r *http.Request) {
	var repair models.Repair
	if err := json.NewDecoder(r.Body).Decode(&repair); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repair.CreatedAt = "2023-10-01"
	repair.UpdatedAt = "2023-10-01"
	result := rt.db.Create(&repair)
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

func (rt *router) updateRepair(w http.ResponseWriter, r *http.Request) {
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
	result := rt.db.Model(&models.Repair{}).Where("id = ?", repairID).Updates(repair)
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

func (rt *router) deleteRepair(w http.ResponseWriter, r *http.Request) {
	repairID := r.URL.Path[len("/repairs/"):]
	if repairID == "" {
		http.Error(w, "Repair ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.Repair{}, repairID)
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

// RepairStatus CRUD handlers
func (rt *router) getRepairStatus(w http.ResponseWriter, r *http.Request) {
	var repairStatuses []models.RepairStatus
	result := rt.db.Find(&repairStatuses)
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

func (rt *router) inputRepairStatus(w http.ResponseWriter, r *http.Request) {
	var repairStatus models.RepairStatus
	if err := json.NewDecoder(r.Body).Decode(&repairStatus); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repairStatus.UpdatedAt = "2023-10-01"
	result := rt.db.Create(&repairStatus)
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

func (rt *router) updateRepairStatus(w http.ResponseWriter, r *http.Request) {
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
	result := rt.db.Model(&models.RepairStatus{}).Where("id = ?", statusID).Updates(repairStatus)
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

func (rt *router) deleteRepairStatus(w http.ResponseWriter, r *http.Request) {
	statusID := r.URL.Path[len("/repair-statuses/"):]
	if statusID == "" {
		http.Error(w, "Repair Status ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.RepairStatus{}, statusID)
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

// ProductUpdateHistory CRUD handlers
func (rt *router) getProductUpdateHistory(w http.ResponseWriter, r *http.Request) {
	var histories []models.ProductUpdateHistory
	result := rt.db.Find(&histories)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve product update histories", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   histories,
		"status": "success",
		"count":  len(histories),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) inputProductUpdateHistory(w http.ResponseWriter, r *http.Request) {
	var history models.ProductUpdateHistory
	if err := json.NewDecoder(r.Body).Decode(&history); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	history.UpdatedAt = "2023-10-01"
	result := rt.db.Create(&history)
	if result.Error != nil {
		http.Error(w, "Failed to create product update history", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product update history created successfully",
		"status":  "success",
		"data":    history,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) deleteProductUpdateHistory(w http.ResponseWriter, r *http.Request) {
	historyID := r.URL.Path[len("/product-histories/"):]
	if historyID == "" {
		http.Error(w, "Product Update History ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.ProductUpdateHistory{}, historyID)
	if result.Error != nil {
		http.Error(w, "Failed to delete product update history", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Product update history not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Product update history deleted successfully",
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

// Payment CRUD handlers
func (rt *router) getPayment(w http.ResponseWriter, r *http.Request) {
	var payments []models.Payment
	result := rt.db.Find(&payments)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve payments", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   payments,
		"status": "success",
		"count":  len(payments),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) inputPayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	payment.CreatedAt = "2023-10-01"
	result := rt.db.Create(&payment)
	if result.Error != nil {
		http.Error(w, "Failed to create payment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Payment created successfully",
		"status":  "success",
		"data":    payment,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) updatePayment(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Path[len("/payments/"):]
	if paymentID == "" {
		http.Error(w, "Payment ID is required", http.StatusBadRequest)
		return
	}

	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := rt.db.Model(&models.Payment{}).Where("id = ?", paymentID).Updates(payment)
	if result.Error != nil {
		http.Error(w, "Failed to update payment", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Payment updated successfully",
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

func (rt *router) deletePayment(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Path[len("/payments/"):]
	if paymentID == "" {
		http.Error(w, "Payment ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.Payment{}, paymentID)
	if result.Error != nil {
		http.Error(w, "Failed to delete payment", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Payment deleted successfully",
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

// Shipping CRUD handlers
func (rt *router) getShipping(w http.ResponseWriter, r *http.Request) {
	var shippings []models.Shipping
	result := rt.db.Find(&shippings)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve shippings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data":   shippings,
		"status": "success",
		"count":  len(shippings),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) inputShipping(w http.ResponseWriter, r *http.Request) {
	var shipping models.Shipping
	if err := json.NewDecoder(r.Body).Decode(&shipping); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shipping.CreatedAt = "2023-10-01"
	result := rt.db.Create(&shipping)
	if result.Error != nil {
		http.Error(w, "Failed to create shipping", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Shipping created successfully",
		"status":  "success",
		"data":    shipping,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (rt *router) updateShipping(w http.ResponseWriter, r *http.Request) {
	shippingID := r.URL.Path[len("/shippings/"):]
	if shippingID == "" {
		http.Error(w, "Shipping ID is required", http.StatusBadRequest)
		return
	}

	var shipping models.Shipping
	if err := json.NewDecoder(r.Body).Decode(&shipping); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := rt.db.Model(&models.Shipping{}).Where("id = ?", shippingID).Updates(shipping)
	if result.Error != nil {
		http.Error(w, "Failed to update shipping", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Shipping not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Shipping updated successfully",
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

func (rt *router) deleteShipping(w http.ResponseWriter, r *http.Request) {
	shippingID := r.URL.Path[len("/shippings/"):]
	if shippingID == "" {
		http.Error(w, "Shipping ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.Shipping{}, shippingID)
	if result.Error != nil {
		http.Error(w, "Failed to delete shipping", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Shipping not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Shipping deleted successfully",
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

// ProductPerOrder CRUD handlers
func (rt *router) getProductPerOrder(w http.ResponseWriter, r *http.Request) {
	var productOrders []models.ProductPerOrder
	result := rt.db.Find(&productOrders)
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

func (rt *router) inputProductPerOrder(w http.ResponseWriter, r *http.Request) {
	var productOrder models.ProductPerOrder
	if err := json.NewDecoder(r.Body).Decode(&productOrder); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	productOrder.CreatedAt = "2023-10-01"
	result := rt.db.Create(&productOrder)
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

func (rt *router) deleteProductPerOrder(w http.ResponseWriter, r *http.Request) {
	productOrderID := r.URL.Path[len("/product-orders/"):]
	if productOrderID == "" {
		http.Error(w, "Product Order ID is required", http.StatusBadRequest)
		return
	}

	result := rt.db.Delete(&models.ProductPerOrder{}, productOrderID)
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
