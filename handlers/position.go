package handlers

import (
	"amx_support_api/cache"
	constants "amx_support_api/constants"
	mapper "amx_support_api/mappers"
	"amx_support_api/models"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	// "bytes"
	// "encoding/binary"
	"net/http"

	"github.com/gin-gonic/gin"
)


func MockPositionController(c *gin.Context) {
	req := &models.PositionReportReq{}
	if err := c.BindJSON(req); err != nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//login for access token
	accessToken := LoginBearerToken()	
    fmt.Println("ac",accessToken)
	//get secInfo
	getSecInfo,err:= GetSecInfo(req,accessToken)
	fmt.Println(getSecInfo)
	if err != nil {
        fmt.Println("Error sending request:", err)
        return 
    }
	redisClient,err:= cache.GetPositionCacheClient()
	if err != nil || redisClient == nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	res := &models.PositionReportResp{} 
    mapper.MapPositionJsonToByte(req, res, getSecInfo)

	buf := &bytes.Buffer{}
	errs := binary.Write(buf, binary.LittleEndian, res)
	if errs != nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	key := string(bytes.Trim(res.StrUserID[:], "\x00")) + "|POSITION"
	field := string(bytes.Trim(res.StrSymbol[:], "\x00")) + "|" + req.StrProduct

	err = redisClient.Del(key).Err()
	if err != nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//redis set
	fmt.Println("buf",buf.Bytes());
	err = redisClient.HSet(key, field, buf.Bytes()).Err()
	if err != nil {
		response := models.APIRes{
			Message:   "FAILURE",
			ErrorCode: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := models.APIRes{
		Status:    true,
		Message:   "SUCCESS",
		ErrorCode: "",
		Data:      req,
	}
	c.JSON(http.StatusOK, response)
}

func GetSecInfo(req *models.PositionReportReq, accessToken string)(map[string]interface{},error){
	apiURL := constants.GetSecInfoEndPoint
	payload := struct {
        Key1 string `json:"exchange"`
        Key2 string   `json:"symbol"`
    }{
        Key1: req.StrExchSegment,
        Key2: req.StrSymbol,
    }

	payloadBytes, err := json.Marshal(payload)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return nil, errors.New("An error occurred 1")
    }
	secreq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return nil, errors.New("An error occurred 2")
    }
    secreq.Header.Set("Authorization", "Bearer "+ accessToken)

	client := &http.Client{}
    response, err := client.Do(secreq)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, errors.New("An error occurred 3")
    }
    defer response.Body.Close()
    if response.StatusCode != http.StatusOK {
        fmt.Println("Error: Unexpected status code", response.StatusCode)
        return nil, errors.New("An error occurred 4")
	}
    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return nil, errors.New("An error occurred 5")
    }
	var responseData map[string]interface{}
    if err := json.Unmarshal(responseBody, &responseData); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return nil, errors.New("An error occurred 6")
    }
	dataMap := responseData["data"].(map[string]interface{})
	return dataMap,nil
}

func LoginBearerToken()(string){
	apiURL := constants.LoginEndPoint
	payload := struct {
        Key1 string `json:"requestIDType"`
        Key2 string   `json:"requestID"`
		Key3 string  `json:"passOrPIN"`
    }{
        Key1: "USERID",
        Key2: "AUTOTEST1",
		Key3: "1111",
    }
	payloadBytes, err := json.Marshal(payload)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return err.Error()
    }
	loginreq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return err.Error()
    }
	loginreq.Header.Set("X-SourceID", "5")
	loginreq.Header.Set("X-UserType", "1")
	loginreq.Header.Set("X-OperatingSystem", "Ubuntu")
	loginreq.Header.Set("X-DeviceID", "1234")
    loginreq.Header.Set("Authorization", "Bearer "+constants.LoginBearerToken)

	client := &http.Client{}
    response, err := client.Do(loginreq)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return err.Error()
    }
    defer response.Body.Close()
    if response.StatusCode != http.StatusOK {
        fmt.Println("Error: Unexpected status code", response.StatusCode)
        return err.Error()
    }
    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return err.Error()
    }
	var responseData map[string]interface{}
    if err := json.Unmarshal(responseBody, &responseData); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return err.Error()
    }
	dataMap := responseData["data"].(map[string]interface{})
	accessToken := dataMap["accesstoken"].(string)
	return accessToken
}
