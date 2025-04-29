package main

import (
	"fmt"
	"log"
	"net/http"
	"soulstreet/internal/config/cors"
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
	
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /product", handler.CreateProduct)
	mux.HandleFunc("GET /products", handler.GetAll)
	mux.HandleFunc("GET /product", handler.GetByID)
	mux.Handle("GET /images/", http.StripPrefix("/images/", http.FileServer(http.Dir("uploads"))))

	
	fmt.Println("Servidor rodando http://localhost:8080")
	
	handlerCors := cors.CORS()(mux)
	
	http.ListenAndServe(":8080", handlerCors)
}