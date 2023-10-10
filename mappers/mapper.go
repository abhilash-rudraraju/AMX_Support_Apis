package mapper

import "amx_support_api/models"

// func MapOrderbookJsontoByte(req *models.OrderReportReq, res *models.OrderReportResp) {
// 	res.DStrikePrice = req.DStrikePrice
// 	res.LLotSize = req.LLotSize
// 	res.NPrecision = req.NPrecision
// 	res.NMultiplier = req.NMultiplier
// 	res.DTickSize = req.DTickSize
// 	res.DPrice = req.DPrice
// 	res.DTriggerPrice = req.DTriggerPrice
// 	res.LQuantity = req.LQuantity
// 	res.LDiscQuantity = req.LDiscQuantity
// 	res.LCancelledSize = req.LCancelledSize
// 	res.LClassification = req.LClassification
// 	res.DMktProtectionPrice = req.DMktProtectionPrice
// 	res.DTickUpSqrOff = req.DTickUpSqrOff
// 	res.DTickDownStopLoss = req.DTickDownStopLoss
// 	res.DTrailingTickStopLoss = req.DTrailingTickStopLoss
// 	res.DAvgPrice = req.DAvgPrice
// 	res.LFilledShares = req.LFilledShares
// 	res.LUnfilledSize = req.LUnfilledSize

// 	copy(res.StrUserID[:], []uint8(req.StrUserID))
// 	copy(res.StrExchSeg[:], []uint8(req.StrExchSeg))
// 	copy(res.StrAccountID[:], []uint8(req.StrAccountID))
// 	copy(res.StrOrdDuration[:], []uint8(req.StrOrdDuration))
// 	copy(res.StrCustomerFirm[:], []uint8(req.StrCustomerFirm))
// 	copy(res.StrProduct[:], []uint8(req.StrProduct))
// 	copy(res.StrOrderType[:], []uint8(req.StrOrderType))
// 	copy(res.StrTrdSymbol[:], []uint8(req.StrTrdSymbol))
// 	copy(res.StrTransType[:], []uint8(req.StrTransType))
// 	copy(res.StrGuiOrdID[:], []uint8(req.StrGuiOrdID))
// 	copy(res.StrExecBroker[:], []uint8(req.StrExecBroker))
// 	copy(res.StrGuiOrgOrdID[:], []uint8(req.StrGuiOrgOrdID))
// 	copy(res.StrStrategyCode[:], []uint8(req.StrStrategyCode))
// 	copy(res.StrBasketID[:], []uint8(req.StrBasketID))
// 	copy(res.StrUniqueKey[:], []uint8(req.StrUniqueKey))
// 	copy(res.StrToken[:], []uint8(req.StrToken))
// 	copy(res.StrInstType[:], []uint8(req.StrToken))
// 	copy(res.StrGroup[:], []uint8(req.StrGroup))
// 	copy(res.StrExpiryDate[:], []uint8(req.StrExpiryDate))
// 	copy(res.StrOptionType[:], []uint8(req.StrOptionType))
// 	copy(res.StrSymbolName[:], []uint8(req.StrSymbolName))
// 	copy(res.StrSymbolDesc[:], []uint8(req.StrSymbolDesc))
// 	copy(res.StrStrategyID[:], []uint8(req.StrStrategyID))
// 	copy(res.StrStrTrgSeqNo[:], []uint8(req.StrStrTrgSeqNo))
// 	copy(res.StrUser[:], []uint8(req.StrUser))
// 	copy(res.StrLtpOrAtp[:], []uint8(req.StrLtpOrAtp))
// 	copy(res.StrSqrOffAbsOrTick[:], []uint8(req.StrSqrOffAbsOrTick))
// 	copy(res.StrSLAbsOrTick[:], []uint8(req.StrSLAbsOrTick))
// 	copy(res.StrTrailTickYesNo[:], []uint8(req.StrTrailTickYesNo))
// 	copy(res.StrVendorCode[:], []uint8(req.StrVendorCode))
// 	copy(res.StrNestOrdNum[:], []uint8(req.StrNestOrdNum))
// 	copy(res.StrReqID[:], []uint8(req.StrReqID))
// 	copy(res.StrExchOrdID[:], []uint8(req.StrExchOrdID))
// 	copy(res.StrText[:], []uint8(req.StrText))
// 	copy(res.StrStatus[:], []uint8(req.StrStatus))
// 	copy(res.StrOrdStatus[:], []uint8(req.StrOrdStatus))
// 	copy(res.StrNestUpdateTime[:], []uint8(req.StrNestUpdateTime))
// 	copy(res.StrExchTime[:], []uint8(req.StrExchTime))
// 	copy(res.StrExchOrdUpdateTime[:], []uint8(req.StrExchOrdUpdateTime))
// 	copy(res.StrRejectionBy[:], []uint8(req.StrRejectionBy))
// 	copy(res.StrOrderGenType[:], []uint8(req.StrOrderGenType))
// 	copy(res.StrSyomOrderID[:], []uint8(req.StrSyomOrderID))
// 	copy(res.StrFillID[:], []uint8(req.StrFillID))
// 	copy(res.StrSymbol[:], []uint8(req.StrSymbol))
// 	copy(res.StrFillTime[:], []uint8(req.StrFillTime))
// 	copy(res.StrOrdValidityDate[:], []uint8(req.StrOrdValidityDate))
// }

func MapPositionJsonToByte(req *models.PositionReportReq, res *models.PositionReportResp) {
	res.DStrikePrice = req.DStrikePrice
	res.LLotSize = req.LLotSize
	res.NPrecision = req.NPrecision
	res.NMultiplier = req.NMultiplier
	res.DTickSize = req.DTickSize
	res.DPriceNum = req.DPriceNum
	res.DGenDen = req.DGenDen
	res.DGenNum = req.DGenNum
	res.DPriceDen = req.DPriceDen
	res.IBoardLotQty = req.IBoardLotQty
	res.DBuyQty = req.DBuyQty
	res.DSellQty = req.DSellQty
	res.DBuyAmount = req.DBuyAmount
	res.DSellAmount = req.DSellAmount
	res.DCfBuyQty = req.DCfBuyQty
	res.DCfSellQty = req.DCfSellQty
	res.DCfBuyAmount = req.DCfBuyAmount
	res.DCfSellAmount = req.DCfSellAmount
	res.DBuyAvgPrice = req.DBuyAvgPrice
	res.DSellAvgPrice = req.DSellAvgPrice
	res.DAvgNetPrice = req.DAvgNetPrice
	res.DNetValue = req.DNetValue
	res.DNetQty = req.DNetQty
	res.DTotalBuyValue = req.DTotalBuyValue
	res.DTotalSellValue = req.DTotalSellValue
	res.DCfBuyAvgPrice = req.DCfBuyAvgPrice
	res.DCfSellAvgPrice = req.DCfSellAvgPrice
	res.DTotalBuyAvgPrice = req.DTotalBuyAvgPrice
	res.DTotalSellAvgPrice = req.DTotalSellAvgPrice
	res.DNetPrice = req.DNetPrice

	copy(res.StrUserID[:], []uint8(req.StrUserID))
	copy(res.StrExchSegment[:], []uint8(req.StrExchSegment))
	copy(res.StrSymbol[:], []uint8(req.StrSymbol))
	copy(res.StrAccountID[:], []uint8(req.StrAccountID))
	copy(res.StrProduct[:], []uint8(req.StrProduct))
	copy(res.StrTrdSymbol[:], []uint8(req.StrTrdSymbol))
	copy(res.StrInstType[:], []uint8(req.StrInstType))
	copy(res.StrGroup[:], []uint8(req.StrGroup))
	copy(res.StrExpiryDate[:], []uint8(req.StrExpiryDate))
	copy(res.StrOptionType[:], []uint8(req.StrOptionType))
	copy(res.StrSymbolName[:], []uint8(req.StrSymbolName))
	copy(res.StrSymbolDesc[:], []uint8(req.StrSymbolDesc))
}

// func MapTradeBookJsonToByte(req *models.TradeReportReq, res *models.TradeReportResp) {
// 	res.DStrikePrice = req.DStrikePrice
// 	res.LLotSize = req.LLotSize
// 	res.NPrecision = req.NPrecision
// 	res.NMultiplier = req.NMultiplier
// 	res.DTickSize = req.DTickSize
// 	res.DFillPrice = req.DFillPrice
// 	res.LFillSize = req.LFillSize
// 	res.SiFillLeg = req.SiFillLeg

// 	copy(res.StrUserID[:], []uint8(req.StrUserID))
// 	copy(res.StrAccountID[:], []uint8(req.StrAccountID))
// 	copy(res.StrBasketID[:], []uint8(req.StrBasketID))
// 	copy(res.StrCustomerFirm[:], []uint8(req.StrCustomerFirm))
// 	copy(res.StrExchOrdID[:], []uint8(req.StrExchOrdID))
// 	copy(res.StrExchSeg[:], []uint8(req.StrExchSeg))
// 	copy(res.StrExchTime[:], []uint8(req.StrExchTime))
// 	copy(res.StrFillDate[:], []uint8(req.StrFillDate))
// 	copy(res.StrFillID[:], []uint8(req.StrFillID))
// 	copy(res.StrFillStatus[:], []uint8(req.StrFillStatus))
// 	copy(res.StrFillTime[:], []uint8(req.StrFillTime))
// 	copy(res.StrFromAccID[:], []uint8(req.StrFromAccID))
// 	copy(res.StrGuiOrdID[:], []uint8(req.StrGuiOrdID))
// 	copy(res.StrGuiOrgOrdID[:], []uint8(req.StrGuiOrgOrdID))
// 	copy(res.StrLineID[:], []uint8(req.StrLineID))
// 	copy(res.StrNestOrdNum[:], []uint8(req.StrNestOrdNum))
// 	copy(res.StrNestTimeMiliSec[:], []uint8(req.StrNestTimeMiliSec))
// 	copy(res.StrNestTimeSec[:], []uint8(req.StrNestTimeSec))
// 	copy(res.StrOrderGenType[:], []uint8(req.StrOrderGenType))
// 	copy(res.StrPriceType[:], []uint8(req.StrPriceType))
// 	copy(res.StrProduct[:], []uint8(req.StrProduct))
// 	copy(res.StrReportType[:], []uint8(req.StrReportType))
// 	copy(res.StrStatus[:], []uint8(req.StrStatus))
// 	copy(res.StrStrategyCode[:], []uint8(req.StrStrategyCode))
// 	copy(res.StrSymbol[:], []uint8(req.StrSymbol))
// 	copy(res.StrInstType[:], []uint8(req.StrInstType))
// 	copy(res.StrGroup[:], []uint8(req.StrGroup))
// 	copy(res.StrExpiryDate[:], []uint8(req.StrExpiryDate))
// 	copy(res.StrOptionType[:], []uint8(req.StrOptionType))
// 	copy(res.StrSymbolName[:], []uint8(req.StrSymbolName))
// 	copy(res.StrSymbolDesc[:], []uint8(req.StrSymbolDesc))
// 	copy(res.StrSyomOrderID[:], []uint8(req.StrSyomOrderID))
// 	copy(res.StrTransType[:], []uint8(req.StrTransType))
// 	copy(res.StrTrdSymbol[:], []uint8(req.StrTrdSymbol))
// 	copy(res.StrUser[:], []uint8(req.StrUser))
// 	copy(res.StrWarehouseLoc[:], []uint8(req.StrWarehouseLoc))
// 	copy(res.StrWaveID[:], []uint8(req.StrWaveID))
// 	copy(res.StrModifiedBy[:], []uint8(req.StrModifiedBy))
// 	copy(res.StrStrategyID[:], []uint8(req.StrStrategyID))
// 	copy(res.StrStrategyCategory[:], []uint8(req.StrStrategyCategory))
// 	copy(res.StrPan[:], []uint8(req.StrPan))
// 	copy(res.StrExecBroker[:], []uint8(req.StrExecBroker))
// }

// func MapRmsLimitEntityReqJsonToResJson(req *models.AMXGetRmsLimitEntityReq, res *models.AMXGetRmsLimitEntityResp) {
// 	res.BoardLotLimit = req.BoardLotLimit
// 	res.DCollateralValue = req.DCollateralValue
// 	res.DCollateral = req.DCollateral
// 	res.DRmsCollateral = req.DRmsCollateral
// 	res.DAdhocMargin = req.DAdhocMargin
// 	res.DNotionalCash = req.DNotionalCash
// 	res.DCdsSpreadBenefit = req.DCdsSpreadBenefit
// 	res.DAmountUtilizedPrsnt = req.DAmountUtilizedPrsnt
// 	res.DUnrealizedMtomPrsnt = req.DUnrealizedMtomPrsnt
// 	res.DRealizedMtomPrsnt = req.DRealizedMtomPrsnt
// 	res.DExposureMarginPrsnt = req.DExposureMarginPrsnt
// 	res.DSpanMarginPrsnt = req.DSpanMarginPrsnt
// 	res.DNfospreadBenefit = req.DNfospreadBenefit
// 	res.DPremiumPrsnt = req.DPremiumPrsnt
// 	res.DMarginVarPrsnt = req.DMarginVarPrsnt
// 	res.DMarginScripBasketPrsnt = req.DMarginScripBasketPrsnt
// 	res.DMtomWarningPrcntPrsnt = req.DMtomWarningPrcntPrsnt
// 	res.DMarginWarningPrcntPrsnt = req.DMarginWarningPrcntPrsnt
// 	res.DMtomSquareOffWarningPrcntPrsnt = req.DMtomSquareOffWarningPrcntPrsnt
// 	res.DMarginUsed = req.DMarginUsed
// 	res.DRmsPayInAmt = req.DRmsPayInAmt
// 	res.DRmsPayOutAmt = req.DRmsPayOutAmt
// 	res.DNet = req.DNet
// 	res.DSpecialMarginPrsnt = req.DSpecialMarginPrsnt
// 	res.DDaneLimit = req.DDaneLimit
// 	res.DDeliveryMarginPresent = req.DDeliveryMarginPresent
// 	res.DCncSellcrdPresent = req.DCncSellcrdPresent
// 	res.DCollateralValueMult = req.DCollateralValueMult
// 	res.DAdhocLimitMult = req.DAdhocLimitMult
// 	res.DRmsCollateralMult = req.DRmsCollateralMult
// 	res.DNationalCashMult = req.DNationalCashMult
// 	res.DComSpanMrgnNrmlPrsnt = req.DComSpanMrgnNrmlPrsnt
// 	res.DComSpanMrgnMisPrsnt = req.DComSpanMrgnMisPrsnt
// 	res.DComExpsrMrgnNrmlPrsnt = req.DComExpsrMrgnNrmlPrsnt
// 	res.DComExpsrMrgnMisPrsnt = req.DComExpsrMrgnMisPrsnt
// 	res.DAddMrgnNrmlPrsnt = req.DAddMrgnNrmlPrsnt
// 	res.DAddMrgnMisPrsnt = req.DAddMrgnMisPrsnt
// 	res.DAddPreExpMrgnNrmlPrsnt = req.DAddPreExpMrgnNrmlPrsnt
// 	res.DAddPreExpMrgnMisPrsnt = req.DAddPreExpMrgnMisPrsnt
// 	res.DSplMrgnNrmlPrsnt = req.DSplMrgnNrmlPrsnt
// 	res.DSplMrgnMisPrsnt = req.DSplMrgnMisPrsnt
// 	res.DTenderMrgnNrmlPrsnt = req.DTenderMrgnNrmlPrsnt
// 	res.DTenderMrgnMisPrsnt = req.DTenderMrgnMisPrsnt
// 	res.DDeliveryMrgnNrmlPrsnt = req.DDeliveryMrgnNrmlPrsnt
// 	res.DDeliveryMrgnMisPrsnt = req.DDeliveryMrgnMisPrsnt
// 	res.DCurPremiumNrmlPrsnt = req.DCurPremiumNrmlPrsnt
// 	res.DCurExpMrgnNrmlPrsnt = req.DCurExpMrgnNrmlPrsnt
// 	res.DCurExpMrgnMisPrsnt = req.DCurExpMrgnMisPrsnt
// 	res.DCurPremiumMisPrsnt = req.DCurPremiumMisPrsnt
// 	res.DCurSpanMrgnNrmlPrsnt = req.DCurSpanMrgnNrmlPrsnt
// 	res.DCurSpanMrgnMisPrsnt = req.DCurSpanMrgnMisPrsnt
// 	res.DFoPremiumNrmlPrsnt = req.DFoPremiumNrmlPrsnt
// 	res.DFoPremiumMisPrsnt = req.DFoPremiumMisPrsnt
// 	res.DFoExpMrgnNrmlPrsnt = req.DFoExpMrgnNrmlPrsnt
// 	res.DFoExpMrgnMisPrsnt = req.DFoExpMrgnMisPrsnt
// 	res.DFoSpanrgnNrmlPrsnt = req.DFoSpanrgnNrmlPrsnt
// 	res.DFoSpanrgnMisPrsnt = req.DFoSpanrgnMisPrsnt
// 	res.DMrgnScrpBsktNrmlPrsnt = req.DMrgnScrpBsktNrmlPrsnt
// 	res.DMrgnScrpBsktMisPrsnt = req.DMrgnScrpBsktMisPrsnt
// 	res.DMrgnScrpBsktCncPrsnt = req.DMrgnScrpBsktCncPrsnt
// 	res.DMrgnVarNrmlPrsnt = req.DMrgnVarNrmlPrsnt
// 	res.DMrgnVarMisPrsnt = req.DMrgnVarMisPrsnt
// 	res.DAmtUntilizedPrsnt = req.DAmtUntilizedPrsnt
// 	res.DCncMrgnVarPrsnt = req.DCncMrgnVarPrsnt
// 	res.DMarginUsedPrsnt = req.DMarginUsedPrsnt
// 	res.DCashUnRlsMtomPrsnt = req.DCashUnRlsMtomPrsnt
// 	res.DFoUnRlsMtomPrsnt = req.DFoUnRlsMtomPrsnt
// 	res.DComUnRlsMtomPrsnt = req.DComUnRlsMtomPrsnt
// 	res.DCurUnRlsMtomPrsnt = req.DCurUnRlsMtomPrsnt
// 	res.DCashRlsMtomPrsnt = req.DCashRlsMtomPrsnt
// 	res.DFoRlsMtomPrsnt = req.DFoRlsMtomPrsnt
// 	res.DComRlsMtomPrsnt = req.DComRlsMtomPrsnt
// 	res.DCurRlsMtomPrsnt = req.DCurRlsMtomPrsnt

// 	copy(res.StrCategory[:], []uint8(req.StrCategory))
// 	copy(res.StrExchSeg[:], []uint8(req.StrExchSeg))
// 	copy(res.StrProduct[:], []uint8(req.StrProduct))
// 	copy(res.StrSegment[:], []uint8(req.StrSegment))
// 	copy(res.StrEntity[:], []uint8(req.StrEntity))
// 	copy(res.StrAccountId[:], []uint8(req.StrAccountId))

// }

// func MapMarginJsonToByte(req *models.AMXClientProfitReq, res *models.AMXClientProfitRes) {
// 	res.SqOffAmt = req.SqOffAmt

// 	copy(res.ClientCode[:], []uint8(req.ClientCode))
// }

// func MapDelieveryJsonToByte(req *models.AMXDeliveryValueReq, res *models.AMXDeliveryValueRes) {
// 	res.DBuyAmount = req.DBuyAmount
// 	res.DSellAmount = req.DSellAmount

// 	copy(res.StrExchSegment[:], []uint8(req.StrExchSegment))
// 	copy(res.StrAccountId[:], []uint8(req.StrAccountId))
// 	copy(res.StrSymbol[:], []uint8(req.StrSymbol))
// 	copy(res.StrTrdSymbol[:], []uint8(req.StrTrdSymbol))
// }