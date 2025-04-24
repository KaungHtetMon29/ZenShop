package main

import (
	"fmt"
	"go_boilerplate/internal/db_utils"
	"go_boilerplate/internal/models"
	"os"

	"go_boilerplate/internal/routes"
	"go_boilerplate/internal/services"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Env loaded")

	dbConfig := db_utils.NewDBConfig(db_utils.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PW"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable"})

	db, err := dbConfig.ConnectDB(dbConfig.GetDSNWithTimeZone("Asia/Shanghai"))
	services.NewService(db)
	fmt.Println("Connected to database")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	// Migrate the schema
	err = db.AutoMigrate(
		&models.Brand{},
		&models.Category{},
		&models.Product{},
		&models.Repair{},
		&models.RepairStatus{},
		&models.Order{},
		&models.Shipping{},
		&models.ProductPerOrder{},
		&models.Payment{},
		&models.ProductUpdateHistory{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database migrated")

	// Initialize the router
	router := routes.InitializeRoutes(db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins, or specify like []string{"http://localhost:3000"}
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Apply CORS middleware to the router
	handler := corsMiddleware.Handler(router)
	server := http.Server{Addr: ":8080", Handler: handler}

	fmt.Println("Starting server on port 8080")
	fmt.Println("🔥 main() started")
	if err := server.ListenAndServe(); err != nil {
		panic("failed to start server")
	}
}
