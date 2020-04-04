# CyptoServer

This project exposes a web service which creates a micro-service with the following endpoint

  - /currency/{symbol}
    Returns the real-time crypto prices of the given currency symbol.
    Sample Response:
    ```js
    {
        "id": "ETH",
        "fullName": "Ethereum",
        "ask": "0.054464",
        "bid": "0.054463",
        "last": "0.054463",
        "open": "0.057133",
        "low": "0.053615",
        "high": "0.057559",
        "feeCurrency": "BTC"
    }
    ```
  - /currency/all
    Returns the real-time crypto prices of all the supported currencies.
    Sample Response:
    ```js
    {
        "currencies": [
            {
                "id": "ETH",
                "fullName": "Ethereum",
                "ask": "0.054464",
                "bid": "0.054463",
                "last": "0.054463",
                "open": "0.057133",
                "low": "0.053615",
                "high": "0.057559",
                "feeCurrency": "BTC"
            },
            {
                "id": "BTC",
                "fullName": "Bitcoin",
                "ask":"7906.72",
                "bid":"7906.28",
                "last":"7906.48",
                "open":"7952.3",
                "low":"7561.51",
                "high":"8107.96",
                "feeCurrency": "USD"
            }
        ]
    }
    ```

# Powered by API
    - https://api.hitbtc.com/
    - Ticker = "https://api.hitbtc.com/api/2/public/ticker/"
	- Symbol = "https://api.hitbtc.com/api/2/public/symbol/"
	- Currency ="https://api.hitbtc.com/api/2/public/currency/"
    
