package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

// DB - глобальная переменная для хранения соединения с базой данных
var DB *sql.DB

// InitDB инициализирует соединение с базой данных, используя переданную строку подключения
func InitDB(dataSourceName string) {
	var err error
	// Открываем соединение с базой данных
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err) // Если произошла ошибка при открытии соединения, завершаем работу с паникой
	}

	// Проверяем соединение с базой данных
	if err = DB.Ping(); err != nil {
		log.Panic(err) // Если произошла ошибка при проверке соединения, завершаем работу с паникой
	}
}
