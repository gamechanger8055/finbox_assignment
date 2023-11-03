package clients

import (
	"awesomeProject/finbox_project/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// APIKey is your API key for the data source (e.g., Alpha Vantage).
const APIKey = "xxxxxxx"
const baseURL = "https://www.alphavantage.co/query"
const intradayFunction = "TIME_SERIES_INTRADAY"
const newsSentimentFunction = "NEWS_SENTIMENT"

// FetchIntradayData fetches intraday stock data for a specific symbol.
func FetchIntradayDataFromAlphaVantage(symbol string) ([]models.IntradayData, error) {
	// Adjust the interval, output size, and other parameters as needed.
	interval := "5min"
	outputSize := "full"

	url := fmt.Sprintf("%s?function=%s&symbol=%s&interval=%s&outputsize=%s&apikey=%s",
		baseURL, intradayFunction, symbol, interval, outputSize, APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch intraday data. Status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response from the data source and convert it to IntradayData.
	var result map[string]map[string]map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	intradayData := make([]models.IntradayData, 0)
	for timestamp, values := range result["Time Series (5min)"] {
		t, err := time.Parse("2006-01-02 15:04:05", timestamp)
		if err != nil {
			return nil, err
		}

		openPrice, _ := strconv.ParseFloat(values["1. open"], 64)
		highPrice, _ := strconv.ParseFloat(values["2. high"], 64)
		lowPrice, _ := strconv.ParseFloat(values["3. low"], 64)
		closePrice, _ := strconv.ParseFloat(values["4. close"], 64)
		volume, _ := strconv.Atoi(values["5. volume"])

		intradayData = append(intradayData, models.IntradayData{
			Symbol:     symbol,
			Timestamp:  t,
			OpenPrice:  openPrice,
			HighPrice:  highPrice,
			LowPrice:   lowPrice,
			ClosePrice: closePrice,
			Volume:     volume,
		})
	}

	return intradayData, nil
}

// FetchNewsSentiment fetches news sentiment data for a specific symbol.
func FetchNewsSentimentFromAlphaVantage(symbol string) ([]models.NewsSentiment, error) {
	url := fmt.Sprintf("%s?function=%s&tickers=%s&apikey=%s", baseURL, newsSentimentFunction, symbol, APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch news sentiment data. Status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response from the data source and convert it to NewsSentiment.
	var result map[string]map[string]float64
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	newsSentiment := make([]models.NewsSentiment, 0)
	for timestamp, sentiment := range result[symbol] {
		t, err := time.Parse("2006-01-02 15:04:05", timestamp)
		if err != nil {
			return nil, err
		}

		newsSentiment = append(newsSentiment, models.NewsSentiment{
			Symbol:         symbol,
			SentimentScore: sentiment,
			NewsTimestamp:  t,
		})
	}

	return newsSentiment, nil
}
