package mapper

import (
	"amx_support_api/models"
	"fmt"
)

func MapPositionJsonToByte(req *models.PositionReportReq, res *models.PositionReportResp , getSecInfo map[string]interface{}) {
	res.LLotSize = req.LLotSize
	res.DStrikePrice = getSecInfo["strikePrice"].(float64)
	res.NPrecision = int32(getSecInfo["precision"].(float64))
	res.NMultiplier = int32(getSecInfo["multiplier"].(float64))
	res.DTickSize = getSecInfo["tickSize"].(float64)
	res.DPriceNum = getSecInfo["priceNum"].(float64)
	res.DGenDen = getSecInfo["genDen"].(float64)
	res.DGenNum = getSecInfo["genNum"].(float64)
	res.DPriceDen = getSecInfo["priceDen"].(float64)
	res.IBoardLotQty = int32(getSecInfo["boardLotQty"].(float64))
	buyValue := 0.00
	if req.StrTransactionType == "BUY" || req.StrTransactionType == "buy" {
		buyValue = float64(req.LLotSize)
		}
	fmt.Println("buyVAlue",buyValue);
	res.DBuyQty = buyValue
	sellValue := 0.00
	if req.StrTransactionType == "SELL" || req.StrTransactionType == "sell" {
		sellValue = float64(req.LLotSize)
		}
	res.DSellQty = sellValue
	res.DBuyAmount = buyValue * req.StrStrikePrice
	res.DSellAmount = sellValue * req.StrStrikePrice
	res.DCfBuyQty = buyValue
	res.DCfSellQty = sellValue
	res.DCfBuyAmount = buyValue * req.StrStrikePrice
	res.DCfSellAmount = sellValue * req.StrStrikePrice
	res.DBuyAvgPrice = buyValue * req.StrStrikePrice
	res.DSellAvgPrice = sellValue * req.StrStrikePrice
	avgNetValue := buyValue * req.StrStrikePrice
	if float64(sellValue * req.StrStrikePrice)>float64(buyValue * req.StrStrikePrice) {
		avgNetValue = float64(sellValue * req.StrStrikePrice)
		}
	res.DAvgNetPrice = avgNetValue
	res.DNetValue = avgNetValue
	res.DNetQty = float64(req.LLotSize)
	res.DTotalBuyValue = buyValue * req.StrStrikePrice
	res.DTotalSellValue = sellValue * req.StrStrikePrice
	res.DCfBuyAvgPrice = buyValue * req.StrStrikePrice
	res.DCfSellAvgPrice = sellValue * req.StrStrikePrice
	res.DTotalBuyAvgPrice = buyValue * req.StrStrikePrice
	res.DTotalSellAvgPrice = sellValue * req.StrStrikePrice
	res.DNetPrice = avgNetValue * req.StrStrikePrice
	

	copy(res.StrUserID[:], []uint8(req.StrUserID))
	copy(res.StrExchSegment[:], []uint8(req.StrExchSegment))
	copy(res.StrSymbol[:], []uint8(req.StrSymbol))
	copy(res.StrAccountID[:], []uint8(req.StrUserID))
	copy(res.StrProduct[:], []uint8(req.StrProduct))
	copy(res.StrTrdSymbol[:], []uint8(getSecInfo["trdSymbol"].(string)))
	//instType to be fetched
	copy(res.StrInstType[:], []uint8(""))

	copy(res.StrGroup[:], []uint8(getSecInfo["group"].(string)))
	copy(res.StrExpiryDate[:], []uint8(""))
	copy(res.StrOptionType[:], []uint8(""))
	copy(res.StrSymbolName[:], []uint8(getSecInfo["symbolName"].(string)))
	copy(res.StrSymbolDesc[:], []uint8(getSecInfo["symbolDesc"].(string)))
	
	fmt.Println("res", res)
}