package routes

import (
	"fmt"
	"go_boilerplate/internal/handlers"
	"go_boilerplate/internal/middleware"
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

	// Initialize handlers
	baseHandler := handlers.NewHandler(db)
	brandHandler := handlers.NewBrandHandler(baseHandler)
	categoryHandler := handlers.NewCategoryHandler(baseHandler)
	productHandler := handlers.NewProductHandler(baseHandler)
	orderHandler := handlers.NewOrderHandler(baseHandler)
	repairHandler := handlers.NewRepairHandler(baseHandler)
	checkoutHandler := handlers.NewCheckoutHandler(baseHandler)
	shippingHandler := handlers.NewShippingHandler(baseHandler) // Add shipping handler

	// Apply middleware to handlers

	// Checkout route
	r.AddRoute("POST", "/checkout", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(checkoutHandler.ProcessCheckout))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Shipping routes
	r.AddRoute("GET", "/shipping", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(shippingHandler.GetShippings))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("GET", "/shipping/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(shippingHandler.GetShipping))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/shipping", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(shippingHandler.CreateShipping))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/shipping/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(shippingHandler.UpdateShipping))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/shipping/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(shippingHandler.DeleteShipping))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Brand routes
	r.AddRoute("GET", "/brands", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(brandHandler.GetBrands))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/brands", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(brandHandler.CreateBrand))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/brands/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(brandHandler.DeleteBrand))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/brands/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(brandHandler.UpdateBrand))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Category routes
	r.AddRoute("GET", "/categories", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(categoryHandler.GetCategories))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/categories", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(categoryHandler.CreateCategory))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/categories/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(categoryHandler.DeleteCategory))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/categories/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(categoryHandler.UpdateCategory))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Product routes
	r.AddRoute("GET", "/products", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(productHandler.GetProducts))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("GET", "/products/filter", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(productHandler.GetFilteredProducts))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/products", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(productHandler.CreateProduct))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/products/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(productHandler.UpdateProduct))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/products/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(productHandler.DeleteProduct))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Order routes
	r.AddRoute("GET", "/orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.GetOrders))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.CreateOrder))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/orders/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.UpdateOrder))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/orders/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.DeleteOrder))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// Repair routes
	r.AddRoute("GET", "/repairs", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.GetRepairs))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/repairs", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.CreateRepair))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/repairs/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.UpdateRepair))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/repairs/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.DeleteRepair))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// RepairStatus routes
	r.AddRoute("GET", "/repair-statuses", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.GetRepairStatuses))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/repair-statuses", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.CreateRepairStatus))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("PUT", "/repair-statuses/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.UpdateRepairStatus))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/repair-statuses/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(repairHandler.DeleteRepairStatus))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	// ProductPerOrder routes
	r.AddRoute("GET", "/product-orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.GetProductPerOrders))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("POST", "/product-orders", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.CreateProductPerOrder))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	r.AddRoute("DELETE", "/product-orders/:id", func(w http.ResponseWriter, req *http.Request) {
		handler := middleware.SetHandler(http.HandlerFunc(orderHandler.DeleteProductPerOrder))
		chain := handler.Chain(testmw)
		chain.ServeHTTP(w, req)
	})

	return r
}

func testmw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("🔥 testmw() start")
		// Do middleware processing here

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
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
