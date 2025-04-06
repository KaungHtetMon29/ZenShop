package main

import (
	"fmt"
	"go_boilerplate/internal/db_utils"
	"go_boilerplate/internal/routes"
	"go_boilerplate/internal/services"
	"net/http"
)

func main() {
	dbConfig:= db_utils.NewDBConfig(db_utils.DBConfig{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Password: "12345678",
		DbName: "postgres",
		SSLMode: "disable"})

	db,err:=dbConfig.ConnectDB(dbConfig.GetDSNWithTimeZone("Asia/Shanghai"))
	services.NewService(db);
	
	fmt.Println("Connected to database")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")
	router:=routes.InitializeRoutes()
	server:=http.Server{Addr: ":8080",Handler: router}

	fmt.Println("Starting server on port 8080")
	fmt.Println("🔥 main() started")
	if err := server.ListenAndServe(); err != nil {
		panic("failed to start server")
	}
}