package main

import (
	"fmt"
	"go_boilerplate/internal/routes"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=12345678 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")
	router:= routes.NewRouter()
	router.AddRoute("GET", "/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("GET /users")
	})
	router.AddRoute("GET", "/items", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("GET /users")
	})
	server:=http.Server{Addr: ":8080",Handler: router}
	fmt.Println("Starting server on port 8080")
	fmt.Println("🔥 main() started")
	if err := server.ListenAndServe(); err != nil {
		panic("failed to start server")
	}
}