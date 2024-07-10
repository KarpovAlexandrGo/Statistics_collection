package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stats-service/api"
	"stats-service/db"
	"stats-service/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderBook(t *testing.T) {
	// Initialize test database
	db.InitDB("user=stats_user password=12345 dbname=stats_service sslmode=disable")

	r := gin.Default()
	r.GET("/order_book", api.GetOrderBook)

	req, err := http.NewRequest("GET", "/order_book?exchange_name=test&pair=BTC/USD", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Add more assertions as needed
	var response models.OrderBook
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test", response.Exchange)
	assert.Equal(t, "BTC/USD", response.Pair)
	// Add more assertions for the content of the response
}

func TestSaveOrderBook(t *testing.T) {
	// Initialize test database
	db.InitDB("user=stats_user password=12345 dbname=stats_service sslmode=disable")

	r := gin.Default()
	r.POST("/order_book", api.SaveOrderBook)

	orderBook := models.OrderBook{
		Exchange: "test",
		Pair:     "BTC/USD",
		Asks:     []models.DepthOrder{{Price: 10000, BaseQty: 1}},
		Bids:     []models.DepthOrder{{Price: 9000, BaseQty: 1}},
	}
	jsonValue, err := json.Marshal(orderBook)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/order_book", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Add more assertions as needed
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Order book saved successfully", response["message"])
	// Add more assertions for the content of the response
}
