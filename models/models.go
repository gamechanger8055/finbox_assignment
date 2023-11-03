package models

import "time"

type IntradayData struct {
	ID         int
	Symbol     string
	Timestamp  time.Time
	OpenPrice  float64
	HighPrice  float64
	LowPrice   float64
	ClosePrice float64
	Volume     int
}

type NewsSentiment struct {
	ID                       int
	Symbol                   string
	SentimentScore           float64
	NewsTimestamp            time.Time
	AssociatedIntradayDataID int
}

type AdjustedStockPrice struct {
	ID                     int
	IntradayDataID         int
	ClosingPrice           float64
	SentimentAdjustmentPct float64
	AdjustedClosingPrice   float64
}
