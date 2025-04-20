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
		&models.ProductPerOrder{},
		&models.Category{},
		&models.Product{},
		&models.Brand{},
		&models.ProductUpdateHistory{},
		&models.RepairStatus{},
		&models.Repair{},
		&models.Order{},
		&models.Payment{},
		&models.Shipping{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database migrated")

	// Initialize the router
	router := routes.InitializeRoutes(db)
	server := http.Server{Addr: ":8080", Handler: router}

	fmt.Println("Starting server on port 8080")
	fmt.Println("🔥 main() started")
	if err := server.ListenAndServe(); err != nil {
		panic("failed to start server")
	}
}
