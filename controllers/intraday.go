package controllers

import (
	data_adapters "awesomeProject/finbox_project/data-adapters"
	"awesomeProject/finbox_project/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LiveTickerSimulation(c *gin.Context) {
	symbol := c.Param("symbol")
	timestamp := c.DefaultQuery("timestamp", "")

	intradayData, err := usecases.FetchIntradayData(symbol)
	if err != nil {
		fmt.Errorf("error occured in fetching intraday data", err.Error())
	}
	newsSentiment, err := usecases.FetchNewsSentiment(symbol)
	if err != nil {
		fmt.Errorf("error occured in fetching sentiment data", err.Error())
	}
	adjustedData := usecases.AdjustStockPrices(intradayData, newsSentiment)

	// Store the adjusted data in the database
	database := data_adapters.Connect()
	data_adapters.CreateTables(database)
	data_adapters.StoreData(database, intradayData, newsSentiment, adjustedData)

	// Fetch the adjusted price for the given timestamp
	livePrice := data_adapters.GetLivePrice(database, symbol, timestamp)

	c.JSON(http.StatusOK, gin.H{
		"symbol":    symbol,
		"timestamp": timestamp,
		"livePrice": livePrice,
	})
}
