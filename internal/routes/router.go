package routes

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/internal/middleware"
	"go_boilerplate/internal/models"
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
	addBrandHandler := http.HandlerFunc(r.inputBrand)
	getBrandHandler := http.HandlerFunc(r.getBrand)
	deleteBrandHandler := http.HandlerFunc(r.deleteBrand)
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
	return r
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
