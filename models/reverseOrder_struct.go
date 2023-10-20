package models

type ReverseOrderReq struct {
	Duration         	string  `json:"duration" example:"DAY"`
	Exchange     		string  `json:"exchange" example:"NSE"`
	OrderType         	string  `json:"ordertype" example:"LIMIT"`
	Price         		string  `json:"price" example:"545.00"`
	ProductType         string   `json:"producttype" example:"INTRADAY"`
	Quantity    		string  `json:"quantity" example:"4"`
	Squareoff      		string `json:"squareoff" example:"0.00"`
	Stoploss       		string  `json:"stoploss" example:"0.00"`
	SymbolToken     	string  `json:"symboltoken" example:"3045"`
	TradingSymbol       string  `json:"tradingsymbol" example:"SBIN-EQ"`
	TrailingStopLoss    string  `json:"trailingStopLoss" example:"0.00"`
	TransactionType     string  `json:"transactiontype" example:"BUY"`
	Triggerprice   		string  `json:"triggerprice" example:"0.00"`
	Variety      		string `json:"variety" example:"NORMAL"`
}

type ReverseOrderResponseData struct {
	Script  string `json:"script"`
	OrderID string `json:"orderid"`
}
