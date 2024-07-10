уководство пользователя для проекта "Stats Service"
Введение
"Stats Service" — это сервис для сбора и управления статистическими данными о торговых ордерах. Сервис предоставляет API для получения и сохранения информации о книгах ордеров и истории ордеров.

Установка и запуск
Требования
Go (версия 1.16 или выше)

PostgreSQL (версия 12 или выше)

Установка
Клонируйте репозиторий:
git clone https://github.com/yourusername/stats-service.git
cd stats-service

Установите зависимости:
go mod download

Настройте базу данных:
Создайте базу данных PostgreSQL.
Настройте параметры подключения в файле main.go:
Вставьте свой данные
db.InitDB("user="" password="" dbname="" sslmode=disable")

Запустите сервер:
go run main.go

API Endpoints
Получение книги ордеров
Метод: GET
URL: /order_book

Параметры запроса:

exchange_name (string): Имя биржи.
pair (string): Торговая пара.

Пример запроса:
curl -X GET "http://localhost:8080/order_book?exchange_name=test&pair=BTC/USD"
Пример ответ:
{
  "ID": 1,
  "Exchange": "test",
  "Pair": "BTC/USD",
  "Asks": [
    {
      "Price": 10000,
      "BaseQty": 1
    }
  ],
  "Bids": [
    {
      "Price": 9000,
      "BaseQty": 1
    }
  ]
}
Сохранение книги ордеров
Метод: POST

URL: /order_book

Тело запроса:
{
  "Exchange": "test",
  "Pair": "BTC/USD",
  "Asks": [
    {
      "Price": 10000,
      "BaseQty": 1
    }
  ],
  "Bids": [
    {
      "Price": 9000,
      "BaseQty": 1
    }
  ]
}

Пример запроса:
curl -X POST "http://localhost:8080/order_book" -H "Content-Type: application/json" -d '{
  "Exchange": "test",
  "Pair": "BTC/USD",
  "Asks": [
    {
      "Price": 10000,
      "BaseQty": 1
    }
  ],
  "Bids": [
    {
      "Price": 9000,
      "BaseQty": 1
    }
  ]
}'
Пример ответа:
{
  "message": "Order book saved successfully"
}
Получение истории ордеров
Метод: GET

URL: /order_history

Тело запроса:
{
  "client_name": "test_client"
}
Пример запроса:
curl -X GET "http://localhost:8080/order_history" -H "Content-Type: application/json" -d '{
  "client_name": "test_client"
}'
Пример ответа:
[
  {
    "ClientName": "test_client",
    "ExchangeName": "test_exchange",
    "Label": "test_label",
    "Pair": "BTC/USD",
    "Side": "buy",
    "Type": "market",
    "BaseQty": 1,
    "Price": 10000,
    "AlgorithmNamePlaced": "test_algo",
    "LowestSellPrc": 9900,
    "HighestBuyPrc": 10100,
    "CommissionQuoteQty": 100,
    "TimePlaced": "2023-10-01T12:00:00+03:00"
  }
]
Сохранение истории ордеров
Метод: POST

URL: /order_history

Тело запроса:
{
  "ClientName": "test_client",
  "ExchangeName": "test_exchange",
  "Label": "test_label",
  "Pair": "BTC/USD",
  "Side": "buy",
  "Type": "market",
  "BaseQty": 1,
  "Price": 10000,
  "AlgorithmNamePlaced": "test_algo",
  "LowestSellPrc": 9900,
  "HighestBuyPrc": 10100,
  "CommissionQuoteQty": 100,
  "TimePlaced": "2023-10-01T12:00:00+03:00"
}
Пример запроса:
curl -X POST "http://localhost:8080/order_history" -H "Content-Type: application/json" -d '{
  "ClientName": "test_client",
  "ExchangeName": "test_exchange",
  "Label": "test_label",
  "Pair": "BTC/USD",
  "Side": "buy",
  "Type": "market",
  "BaseQty": 1,
  "Price": 10000,
  "AlgorithmNamePlaced": "test_algo",
  "LowestSellPrc": 9900,
  "HighestBuyPrc": 10100,
  "CommissionQuoteQty": 100,
  "TimePlaced": "2023-10-01T12:00:00+03:00"
}'
Пример ответа:
{
  "message": "Order history saved successfully"
}
Заключение
"Stats Service" предоставляет простой и удобный API для управления статистическими данными о торговых ордерах. Следуя этому руководству, вы сможете легко начать использовать сервис и интегрировать его в свои проекты.

