package main

import (
	"stats-service/api"
	"stats-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных с указанными параметрами
	db.InitDB("user=stats_user password=12345 dbname=stats_service sslmode=disable")

	// Создание экземпляра роутера
	r := gin.Default()

	// Определение маршрутов для API
	r.GET("/order_book", api.GetOrderBook)       // Маршрут для получения книги
	r.POST("/order_book", api.SaveOrderBook)     // Маршрут для сохранения книги
	r.GET("/order_history", api.GetOrderHistory) // Маршрут для получения истории ордеров
	r.POST("/order_history", api.SaveOrder)      // Маршрут для сохранения истории ордеров

	// Запуск сервера
	r.Run()
}
