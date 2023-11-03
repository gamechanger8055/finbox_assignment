### Objective

Retrieve detailed intraday stock data and news sentiment from Alpha Vantage for major tech companies. Adjust stock prices based on sentiment analysis. Design a relational database to store this data efficiently. Develop robust API endpoints to simulate live stock tickers and fetch specific intraday data.

#### Task 1: Retrieve Intraday Data and News Sentiment

- **Data Source**: Use the Alpha Vantage API.

- **Symbols**:
    - AAPL (Apple Inc.)
    - MSFT (Microsoft Corporation)
    - GOOGL (Alphabet Inc.)
    - AMZN (Amazon.com Inc.)
    - TSLA (Tesla Inc.)

- **Endpoints**:
    - **Full Intraday Data**:
      ```
      https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol={SYMBOL}&interval=5min&outputsize=full&apikey=demo
      ```
    - **News Sentiment**:
      ```
      https://www.alphavantage.co/query?function=NEWS_SENTIMENT&tickers={SYMBOL}&apikey=demo
      ```
  Replace `{SYMBOL}` with the respective stock symbol.

- **Guidelines**:
    - Fetch both monthly intraday data and news sentiment for each of the specified stocks.

- **Sentiment Adjustment**:
    - Adjust the replayed stock price on the pre-fetched historical data based on the overall sentiment values from the news sentiment data:
        - **Bearish Sentiment**: Decrease the stock price by 2%.
        - **Somewhat-Bearish Sentiment**: Decrease the stock price by 1%.
        - **Neutral Sentiment**: No adjustment.
        - **Somewhat_Bullish Sentiment**: Increase the stock price by 1%.
        - **Bullish Sentiment**: Increase the stock price by 2%.

#### Task 2: Store and Integrate Data

- **Database Choice**: Opt for a relational database, preferably PostgreSQL. Ensure it's equipped to store intraday stock data, news sentiment, and the sentiment-adjusted closing prices.

- **Table Design**:
    - **Intraday Data**: This table should feature:
        - `ID`: A unique identifier.
        - `Symbol`: Stock symbol (e.g., AAPL).
        - `Timestamp`: The exact date and time of the data point.
        - `OpenPrice`, `HighPrice`, `LowPrice`, `ClosePrice`: OHLC values.
        - `Volume`: Trading volume for that period.

    - **News Sentiment**: This table should encompass:
        - `ID`: A unique identifier.
        - `Symbol`: The associated stock symbol.
        - `SentimentScore`: Numeric representation of sentiment.
        - `NewsTimestamp`: Publication date and time of the news article.
        - `AssociatedIntradayDataID`: A foreign key linking to the `Intraday Data` table, based on a close timestamp match or proximity.

    - **Adjusted Stock Prices**: After data retrieval and sentiment analysis, adjustments will be made to the stock's closing price. This table should feature:
        - `ID`: A unique identifier.
        - `IntradayDataID`: A foreign key linking to `Intraday Data`.
        - `ClosingPrice`: The stock's closing price from the intraday data.
        - `SentimentAdjustmentPercentage`: The adjustment percentage based on sentiment.
        - `AdjustedClosingPrice`: The adjusted closing price after considering the sentiment.

- **Guidelines**:
    - To link the `Intraday Data` and `News Sentiment` tables, use both the stock symbol and timestamp. If a news article was published at 12:00 PM, associate it with intraday data from around that time (e.g., 11:55 AM to 12:05 PM).
    - Uphold data integrity with appropriate constraints and relations between tables.
    - Index vital columns, especially those used in joins or conditions, to ensure efficient querying.

#### Task 3: Develop REST API Endpoints

For this task, you'll be developing one API endpoint to serve the following purpose:

**Live Ticker Simulation**:
- **Endpoint**: `/api/live/{symbol}`
- **HTTP Method**: `GET`
- **Description**: This endpoint should simulate a live ticker price for the specified stock symbol. The price should be replayed from the stored monthly intraday data, based on the same timestamp as the request. This mock live ticker should return the `AdjustedClosingPrice` for the particular timestamp.
- **Response Format**: JSON
- **Response Body Example**:
```json
{
"symbol": "AAPL",
"timestamp": "2023-05-10T12:00:00Z",
"livePrice": 185.23
}
```

Ensure that the API is robust, handles potential errors gracefully, and provides meaningful error messages to the clients.

### Additional Information

- **Alpha Vantage API**: Familiarize yourself with the [Alpha Vantage API](https://www.alphavantage.co/documentation/) to ensure effective data retrieval.

- **Timestamp Linking**: While linking intraday data and sentiment data, focus on `hh:mm:ss` proximity. Dates should be completely disregarded in this process.

- **Simulation Objective**: The challenge's core is not about the exact accuracy of the replayed price. Instead, it emphasizes simulating a live feed using historical intraday data and making sentiment-driven adjustments during the replay. Ensure this simulation is effective and reliable.

Run and build steps.

1. Clone the repo using ssh keys.
2. Connect to db by adding your db credentials.
3. In terminal run the project go run main.go.
4. match the response and db data.
