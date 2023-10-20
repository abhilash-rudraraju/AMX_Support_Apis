package handlers

import (
	constants "amx_support_api/constants"
	"amx_support_api/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)



func ReverserOrderUat(c *gin.Context){
	accessToken:= LoginBearerToken()

	req:= &models.ReverseOrderReq{}
	if err := c.BindJSON(req); err != nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return 
	}
	apiURL := constants.PlaceOrderEndPoint
	payload := struct {
    Duration         	string  `json:"duration" example:"DAY"`
	Exchange     		string  `json:"exchange" example:"NSE"`
	OrderType         	string  `json:"ordertype" example:"LIMIT"`
	Price         		string  `json:"price" example:"545.00"`
	ProductType         string  `json:"producttype" example:"INTRADAY"`
	Quantity    		string  `json:"quantity" example:"4"`
	Squareoff      		string `json:"squareoff" example:"0.00"`
	Stoploss       		string  `json:"stoploss" example:"0.00"`
	SymbolToken     	string  `json:"symboltoken" example:"3045"`
	TradingSymbol       string  `json:"tradingsymbol" example:"SBIN-EQ"`
	TrailingStopLoss    string  `json:"trailingStopLoss" example:"0.00"`
	TransactionType     string  `json:"transactiontype" example:"BUY"`
	Triggerprice   		string  `json:"triggerprice" example:"0.00"`
	Variety      		string `json:"variety" example:"NORMAL"`
    }{
    Duration : req.Duration,
	Exchange : req.Exchange,     		
	OrderType : req.OrderType ,        		
	Price : req.Price,        		
	ProductType : req.ProductType ,       
	Quantity : req.Quantity,   		
	Squareoff : req.Squareoff ,     		
	Stoploss : req.Stoploss ,      		
	SymbolToken : req.SymbolToken ,   	
	TradingSymbol : req.TradingSymbol ,     
	TrailingStopLoss : req.TrailingStopLoss,    
	TransactionType : req.TransactionType,    
	Triggerprice : req.Triggerprice,   		
	Variety : req.Variety ,   		
    }
	payloadBytes, err := json.Marshal(payload)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return 
    }
	placeorderreq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return 
    }
    placeorderreq.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
    response, err := client.Do(placeorderreq)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return 
    }
	if err != nil {
        fmt.Println("Error sending request:", err)
        return 
    }
    defer response.Body.Close()
    if response.StatusCode != http.StatusOK {
        fmt.Println("Error: Unexpected status code", response.StatusCode)
        return 
	}
    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return 
    }
	var responseJSON models.APIRes
    err1 := json.Unmarshal(responseBody, &responseJSON)
    if err1 != nil {
        fmt.Println("Error parsing JSON:", err1)
        return
    }
	//get Order Book
	apiURL=constants.GetOrderBookEndPoint
	getOrderBookReq,err :=http.NewRequest("GET",apiURL,nil)
	if err != nil {
        fmt.Println("Error sending request:", err)
        return 
    }
	getOrderBookReq.Header.Set("Authorization", "Bearer "+accessToken)
	client = &http.Client{}
    response1, err := client.Do(getOrderBookReq)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return 
    }
	defer response1.Body.Close()
    if response1.StatusCode != http.StatusOK {
        fmt.Println("Error: Unexpected status code", response.StatusCode)
        return 
	}
    response1Body, err := ioutil.ReadAll(response1.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return 
    }
	var OrderBookresponseJSON models.APIRes
    err12 := json.Unmarshal(response1Body, &OrderBookresponseJSON)
    if err12 != nil {
        fmt.Println("Error parsing JSON:", err1)
        return
    }
	OrderBookdata:=OrderBookresponseJSON.Data
	PlaceOrderInfo:=responseJSON.Data
	fmt.Println("place",PlaceOrderInfo)
	
     i:=0
	 for i<5 {
	   if checkIfOrderIsComplete(OrderBookdata,PlaceOrderInfo){
		break
	   }
	   time.Sleep(300 * time.Millisecond)
	   i++
	 }
	
	var FinalresponseJSON models.OrderAPIRes
    err13 := json.Unmarshal(responseBody, &FinalresponseJSON)
    if err13 != nil {
        fmt.Println("Error parsing JSON:", err1)
        return
    }
	FinalresponseJSON.OrderStatus = "notcomplete"
	if i!=5 {FinalresponseJSON.OrderStatus = "complete"}
   c.JSON(http.StatusOK, FinalresponseJSON)
}


func checkIfOrderIsComplete(OrderBookdata interface{} ,PlaceOrderInfo interface{})(bool){
	var Orderid string
	if orderInfo, ok := PlaceOrderInfo.(map[string]interface{}); ok {
        orderid, exists := orderInfo["orderid"]
        if exists {
			if orderidString, isString := orderid.(string); isString {
                Orderid = orderidString
                fmt.Printf("Order ID: %v\n", Orderid)
            } else {
                fmt.Println("The 'orderid' key is not of type String.")
            }

            fmt.Printf("Order ID: %v\n", orderid)
        } else {
            fmt.Println("The 'orderid' key does not exist in the map.")
        }
    } else {
        fmt.Println("PlaceOrderInfo is not a map[string]interface{}.")
    }

	if orderBookinfo, ok := OrderBookdata.([]map[string]interface{});ok {
		for _, order := range orderBookinfo {
			orderId, found := order["orderid"].(string)
			if found && orderId==Orderid {
					orderStatus, statusfound := order["status"].(string)
					if statusfound && orderStatus == "complete" {
						return true
			}
		}
	}
}	
return false
}