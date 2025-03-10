package main

import (
	"log"
	"net/http"
	"online-shop/internal/db"
	"online-shop/internal/handlers"
	"online-shop/internal/repository"
	"online-shop/internal/services"
)

func main() {
	db.InitDB()

	productRepo := repository.NewProductRepo(db.DB)
	searchService := services.NewSearchService(productRepo)
	searchHandler := handlers.NewSearchHandler(searchService)

	http.HandleFunc("/search", searchHandler.HandleSearch)

	log.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
