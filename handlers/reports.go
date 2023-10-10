package handlers

import (
	mapper "amx_support_api/mappers"
	"amx_support_api/models"
	// "bytes"
	// "encoding/binary"
	"net/http"

	"github.com/gin-gonic/gin"
)



func MockUpdateMarginController(c *gin.Context) {
	req := &models.PositionReportReq{}
	if err := c.BindJSON(req); err != nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res := &models.PositionReportResp{}
	mapper.MapPositionJsonToByte(req, res)

	// redisClient, err := cache.GetPositionCacheClient()
	// if err != nil || redisClient == nil {
	// 	response := models.APIRes{
	// 		Message:   "FAILURE",
	// 		ErrorCode: err.Error(),
	// 	}
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// buf := &bytes.Buffer{}
	// errs := binary.Write(buf, binary.LittleEndian, res)
	// if errs != nil {
	// 	response := models.APIRes{
	// 		Message:   "FAILURE",
	// 		ErrorCode: err.Error(),
	// 	}
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// key := string(bytes.Trim(res.StrUserID[:], "\x00")) + "|MARGIN"
	// field := string(bytes.Trim(res.StrSymbol[:], "\x00"))

	// err = redisClient.HSet(key, field, buf.Bytes()).Err()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	response := models.APIRes{
		Status:    true,
		Message:   "SUCCESS",
		ErrorCode: "",
		Data:      req,
	}
	c.JSON(http.StatusOK, response)
}