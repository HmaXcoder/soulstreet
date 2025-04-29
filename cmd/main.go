package main

import (
	"fmt"
	"log"
	"net/http"
	"soulstreet/internal/config/db"
	"soulstreet/internal/handler"
	"soulstreet/internal/repository"
	"soulstreet/internal/service"
)

func main() {
	
	dbConn := db.ConnectDB()
	err := db.CreateTableIfNotExist(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewProductRepositoryDB(dbConn)
	service := service.NewProductService(repo)
	handler := handler.NewProductHandler(*service)
	
	http.HandleFunc("POST /products", handler.CreateProduct)
	http.HandleFunc("GET /products", handler.GetAll)
	http.HandleFunc("GET /product", handler.GetByID)

	
	fmt.Println("Servidor rodando http://localhost:8080")
	
	http.ListenAndServe(":8080", nil)
}