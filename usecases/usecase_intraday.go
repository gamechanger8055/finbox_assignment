package usecases

import (
	"awesomeProject/finbox_project/clients"
	"awesomeProject/finbox_project/models"
)

func AdjustStockPrices(intradayData []models.IntradayData, newsSentiment []models.NewsSentiment) []models.AdjustedStockPrice {
	adjustedData := make([]models.AdjustedStockPrice, len(intradayData))
	for i, intraday := range intradayData {
		adjustedClosingPrice := intraday.ClosePrice
		sentiment := newsSentiment[i].SentimentScore

		if sentiment <= -0.02 {
			adjustedClosingPrice *= 0.98
		} else if sentiment <= -0.01 {
			adjustedClosingPrice *= 0.99
		} else if sentiment >= 0.02 {
			adjustedClosingPrice *= 1.02
		} else if sentiment >= 0.01 {
			adjustedClosingPrice *= 1.01
		}

		adjustedData[i] = models.AdjustedStockPrice{
			IntradayDataID:         i + 1,
			ClosingPrice:           intraday.ClosePrice,
			SentimentAdjustmentPct: (adjustedClosingPrice / intraday.ClosePrice) - 1,
			AdjustedClosingPrice:   adjustedClosingPrice,
		}
	}
	return adjustedData
}

func FetchNewsSentiment(symbol string) ([]models.NewsSentiment, error) {
	newsSentiment, err := clients.FetchNewsSentimentFromAlphaVantage(symbol)
	if err != nil {
		return nil, err
	}
	return newsSentiment, nil
}

func FetchIntradayData(symbol string) ([]models.IntradayData, error) {
	intradayData, err := clients.FetchIntradayDataFromAlphaVantage(symbol)
	if err != nil {
		return nil, err
	}
	return intradayData, nil
}
