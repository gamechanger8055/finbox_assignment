package data_adapters

import (
	"awesomeProject/finbox_project/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbUser     = "shubh"
	dbPassword = "1234"
	dbName     = "finbox_db"
)

func Connect() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTables(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS intraday_data (
		id serial PRIMARY KEY,
		symbol text,
		timestamp timestamp,
		open_price real,
		high_price real,
		low_price real,
		close_price real,
		volume int
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS news_sentiment (
		id serial PRIMARY KEY,
		symbol text,
		sentiment_score real,
		news_timestamp timestamp,
		associated_intraday_data_id int
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS adjusted_stock_prices (
		id serial PRIMARY KEY,
		intraday_data_id int,
		closing_price real,
		sentiment_adjustment_pct real,
		adjusted_closing_price real
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func StoreData(db *sql.DB, intradayData []models.IntradayData, newsSentiment []models.NewsSentiment, adjustedData []models.AdjustedStockPrice) {
	// Insert data into tables
	for _, data := range intradayData {
		_, err := db.Exec(`INSERT INTO intraday_data (symbol, timestamp, open_price, high_price, low_price, close_price, volume)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, data.Symbol, data.Timestamp, data.OpenPrice, data.HighPrice, data.LowPrice, data.ClosePrice, data.Volume)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, data := range newsSentiment {
		_, err := db.Exec(`INSERT INTO news_sentiment (symbol, sentiment_score, news_timestamp, associated_intraday_data_id)
			VALUES ($1, $2, $3, $4)`, data.Symbol, data.SentimentScore, data.NewsTimestamp, data.AssociatedIntradayDataID)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, data := range adjustedData {
		_, err := db.Exec(`INSERT INTO adjusted_stock_prices (intraday_data_id, closing_price, sentiment_adjustment_pct, adjusted_closing_price)
			VALUES ($1, $2, $3, $4)`, data.IntradayDataID, data.ClosingPrice, data.SentimentAdjustmentPct, data.AdjustedClosingPrice)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetLivePrice(db *sql.DB, symbol string, timestamp string) float64 {
	// Implement the logic to fetch the adjusted price for the given timestamp
	var livePrice float64
	err := db.QueryRow(`SELECT adjusted_closing_price FROM adjusted_stock_prices
		WHERE intraday_data_id = (SELECT id FROM intraday_data WHERE symbol = $1 AND timestamp = $2)`,
		symbol, timestamp).Scan(&livePrice)
	if err != nil {
		log.Fatal(err)
	}
	return livePrice
}
