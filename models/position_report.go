package models
const (
	ExchSegLen    int = 10
	TokenLen      int = 50
	TrdSymbolLen  int = 50
	UseridLen     int = 21
	AccountidLen  int = 21
	SymbolNameLen int = 50
	ProductLen    int = 10
)

type PositionReportReq struct {
	StrUserID          string  `json:"userid" example:"K111234"`
	StrExchSegment     string  `json:"segment" example:"mcx_fo"`
	StrSymbol          string  `json:"symboltoken" example:"256556"`
	StrProduct         string  `json:"producttype" example:"CARRYFORWARD"`
	LLotSize           int64   `json:"quantity" example:"2"`
	StrTransactionType    string  `json:"transactionType" example:"1"`
	StrStrikePrice      float64 `json:"Price" example:"1"`
}

type PositionReportResp struct {
	StrUserID          [UseridLen]uint8
	StrExchSegment     [ExchSegLen]uint8
	StrSymbol          [SymbolNameLen]uint8
	StrAccountID       [AccountidLen]uint8
	StrProduct         [ProductLen]uint8
	StrTrdSymbol       [TrdSymbolLen]uint8
	StrInstType        [10]uint8
	StrGroup           [15]uint8
	StrExpiryDate      [11]uint8
	DStrikePrice       float64
	StrOptionType      [5]uint8
	StrSymbolName      [SymbolNameLen]uint8
	LLotSize           int64
	NPrecision         int32
	NMultiplier        int32
	StrSymbolDesc      [80]uint8
	DPriceNum          float64
	DGenDen            float64
	DGenNum            float64
	DTickSize          float64
	DPriceDen          float64
	IBoardLotQty       int32
	DBuyQty            float64
	DSellQty           float64
	DBuyAmount         float64
	DSellAmount        float64
	DCfBuyQty          float64
	DCfSellQty         float64
	DCfBuyAmount       float64
	DCfSellAmount      float64
	DBuyAvgPrice       float64
	DSellAvgPrice      float64
	DAvgNetPrice       float64
	DNetValue          float64
	DNetQty            float64 
	DTotalBuyValue     float64
	DTotalSellValue    float64
	DCfBuyAvgPrice     float64
	DCfSellAvgPrice    float64
	DTotalBuyAvgPrice  float64
	DTotalSellAvgPrice float64
	DNetPrice          float64
}