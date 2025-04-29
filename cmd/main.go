package main

import (
	"fmt"
	"net/http"
	"soulstreet/internal/config/db"
	"soulstreet/internal/handler"
	"soulstreet/internal/repository"
	"soulstreet/internal/service"
)

func main() {
	db := db.ConnectDB()
	repo := repository.NewProductRepositoryDB(db)
	service := service.NewProductService(repo)
	handler := handler.NewProductHandler(*service)
	
	http.HandleFunc("/products", handler.CreateProduct)
	
	fmt.Println("Servidor rodando http://localhost:8080")
	
	http.ListenAndServe(":8080", nil)
}