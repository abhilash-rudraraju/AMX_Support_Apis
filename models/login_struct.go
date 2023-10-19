package models


type LoginBody struct{
	IDType    string      `json:"requestIDType"`
	ClientID  string      `json:"requestID"`
	Mpin   string         `json:"passOrPIN"`
}

type LoginResponse struct {
    Status           bool   `json:"status"`
    Message          string `json:"message"`
    ErrorCode        string `json:"errorcode"`
    Data             LoginData   `json:"data"`
}

type LoginData struct {
    UserID               string `json:"userid"`
    Username             string `json:"username"`
    AccessToken          string `json:"accesstoken"`
    RefreshToken         string `json:"refreshtoken"`
    MessageCode          string `json:"messagecode"`
    Message              string `json:"message"`
    LastLoginTime        string `json:"lastlogintime"`
    DaysLeft             int    `json:"daysleft"`
    FeedsTokenID         string `json:"feedstokenid"`
    SubscribedExchProdList string `json:"subscribedexchprodlist"`
    PasswordReset        string `json:"passwordreset"`
    Branch               string `json:"branch"`
    Broker               string `json:"broker"`
    UserType             string `json:"usertype"`
    OrderType            string `json:"ordertype"`
    Exchange             string `json:"exchange"`
    Product              string `json:"product"`
}