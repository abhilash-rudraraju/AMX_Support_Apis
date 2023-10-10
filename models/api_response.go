package models

type APIRes struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"errorcode"`
	Data      interface{} `json:"data"`
}