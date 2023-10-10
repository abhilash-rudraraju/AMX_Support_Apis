package models


type loginBody struct{
	IDType    string      `json:"requestIDType"`
	ClientID  string      `json:"requestID"`
	Mpin   string         `json:"passOrPIN"`
}