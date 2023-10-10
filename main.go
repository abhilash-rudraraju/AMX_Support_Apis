package main

import (
	// "amx_support_apis/handlers"
	// Handler "example/amx_support_api/handlers"
	"amx_support_api/handlers"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	router := gin.Default()
	router.POST("/mockPosition", handlers.MockUpdateMarginController)
	// router.POST("/createPosition",createPosition)

	router.Run("localhost:8020")
}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func createPosition(c *gin.Context){

// 	req:=&models.PositionReportReq{}

// 	if err := c.BindJSON(req); err != nil {
// 		response := models.APIRes{
// 			Message:   "FAILURE",
// 			ErrorCode: err.Error(),
// 		}
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	// res := &models.PositionReportResp{}

// 	// accessToken,err := getAccessToken()

// 	// response, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))

// 	// report.MapPositionJsonToByte(req, res)

// }

// func getAccessToken(){
// 	response, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
// }
