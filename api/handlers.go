package api

import (
	"encoding/json"
	"log"
	"net/http"
	"stats-service/db"
	"stats-service/models"

	"github.com/gin-gonic/gin"
)

// GetOrderBook получает информацию о книге
func GetOrderBook(c *gin.Context) {
	exchangeName := c.Query("exchange_name") // Получаем имя обмена из запроса
	pair := c.Query("pair")                  // Получаем пару из запроса

	// Запрашиваем данные из базы
	rows, err := db.DB.Query("SELECT * FROM order_book WHERE exchange = $1 AND pair = $2", exchangeName, pair)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var orderBook models.OrderBook
	if rows.Next() {
		var asksJSON, bidsJSON []byte
		if err := rows.Scan(
			&orderBook.ID,
			&orderBook.Exchange,
			&orderBook.Pair,
			&asksJSON,
			&bidsJSON,
		); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if err := json.Unmarshal(asksJSON, &orderBook.Asks); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if err := json.Unmarshal(bidsJSON, &orderBook.Bids); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	c.JSON(http.StatusOK, orderBook) // Возвращаем книгу  в формате JSON
}

// SaveOrderBook сохраняет книгу
func SaveOrderBook(c *gin.Context) {
	var orderBook models.OrderBook
	if err := c.BindJSON(&orderBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Преобразуем asks и bids в JSON
	asksJSON, err := json.Marshal(orderBook.Asks)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	bidsJSON, err := json.Marshal(orderBook.Bids)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Вставляем данные в базу
	_, err = db.DB.Exec("INSERT INTO order_book (exchange, pair, asks, bids) VALUES ($1, $2, $3, $4)",
		orderBook.Exchange, orderBook.Pair, asksJSON, bidsJSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order book saved successfully"}) // Возвращаем сообщение об успешном сохранении
}

// GetOrderHistory получает историю ордеров для заданного клиента
func GetOrderHistory(c *gin.Context) {
	var client models.Client
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Запрашиваем данные из базы
	rows, err := db.DB.Query("SELECT * FROM order_history WHERE client_name = $1", client.ClientName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var historyOrders []*models.HistoryOrder
	for rows.Next() {
		var historyOrder models.HistoryOrder
		if err := rows.Scan(
			&historyOrder.ClientName,
			&historyOrder.ExchangeName,
			&historyOrder.Label,
			&historyOrder.Pair,
			&historyOrder.Side,
			&historyOrder.Type,
			&historyOrder.BaseQty,
			&historyOrder.Price,
			&historyOrder.AlgorithmNamePlaced,
			&historyOrder.LowestSellPrc,
			&historyOrder.HighestBuyPrc,
			&historyOrder.CommissionQuoteQty,
			&historyOrder.TimePlaced,
		); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		historyOrders = append(historyOrders, &historyOrder)
	}

	c.JSON(http.StatusOK, historyOrders) // Возвращаем историю ордеров в формате JSON
}

// SaveOrder сохраняет историю ордеров для заданного клиента
func SaveOrder(c *gin.Context) {
	var historyOrder models.HistoryOrder
	if err := c.BindJSON(&historyOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Вставляем данные в базу
	_, err := db.DB.Exec("INSERT INTO order_history (client_name, exchange_name, label, pair, side, type, base_qty, price, algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		historyOrder.ClientName, historyOrder.ExchangeName, historyOrder.Label, historyOrder.Pair, historyOrder.Side, historyOrder.Type, historyOrder.BaseQty, historyOrder.Price, historyOrder.AlgorithmNamePlaced, historyOrder.LowestSellPrc, historyOrder.HighestBuyPrc, historyOrder.CommissionQuoteQty, historyOrder.TimePlaced)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order history saved successfully"}) // Возвращаем сообщение об успешном сохранении
}
