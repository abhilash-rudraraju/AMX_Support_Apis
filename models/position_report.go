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

// type PositionReportReq struct {
// 	StrUserID          string  `json:"userid" example:"K111234"`
// 	StrExchSegment     string  `json:"segment" example:"mcx_fo"`
// 	StrSymbol          string  `json:"symboltoken" example:"256556"`
// 	StrProduct         string  `json:"producttype" example:"CARRYFORWARD"`
// 	LLotSize           int64   `json:"quantity" example:"1"`
// 	TransactionType    string  `json:"type" example:"1"`
// }
type PositionReportReq struct {
	StrUserID          string  `json:"userid" example:"K111234"`
	StrExchSegment     string  `json:"segment" example:"mcx_fo"`
	StrSymbol          string  `json:"symboltoken" example:"256556"`
	StrAccountID       string  `json:"accountid" example:"S111763"`
	StrProduct         string  `json:"producttype" example:"CARRYFORWARD"`
	StrTrdSymbol       string  `json:"tradingsymbol" example:"GOLDPETAL23AUGFUT"`
	StrInstType        string  `json:"instrumenttype" example:"FUTCOM"`
	StrGroup           string  `json:"symbolgroup" example:"XX"`
	StrExpiryDate      string  `json:"expirydate" example:"31AUG2023"`
	DStrikePrice       float64 `json:"strikeprice" example:"0.00"`
	StrOptionType      string  `json:"optiontype" example:"CE"`
	StrSymbolName      string  `json:"symbolname" example:"GOLDPETAL"`
	LLotSize           int64   `json:"lotsize" example:"1"`
	NPrecision         int32   `json:"precision" example:"2"`
	NMultiplier        int32   `json:"multiplier" example:"1"`
	StrSymbolDesc      string  `json:"symbolDesc" example:"2023 Aug 31"`
	DPriceNum          float64 `json:"pricenum" example:"1"`
	DGenDen            float64 `json:"genden" example:"1"`
	DGenNum            float64 `json:"gennum" example:"1"`
	DTickSize          float64 `json:"tickSize" example:"1.00"`
	DPriceDen          float64 `json:"priceden" example:"1"`
	IBoardLotQty       int32   `json:"BoardLotQty" example:"1"`
	DBuyQty            float64 `json:"buyqty" example:"1"`
	DSellQty           float64 `json:"sellqty" example:"1"`
	DBuyAmount         float64 `json:"buyamount" example:"0.00"`
	DSellAmount        float64 `json:"sellamount" example:"0.00"`
	DCfBuyQty          float64 `json:"cfbuyqty" example:"1"`
	DCfSellQty         float64 `json:"cfsellqty" example:"1"`
	DCfBuyAmount       float64 `json:"cfbuyamount" example:"0.00"`
	DCfSellAmount      float64 `json:"cfsellamount" example:"0.00"`
	DBuyAvgPrice       float64 `json:"cfbuyavgprice" example:"0.0"`
	DSellAvgPrice      float64 `json:"cfsellavgprice" example:"0.0"`
	DAvgNetPrice       float64 `json:"avgnetprice" example:"0.0"`
	DNetValue          float64 `json:"netvalue" example:"0.0"`
	DNetQty            float64 `json:"netqty" example:"0"` //determines is squareOff if its zero
	DTotalBuyValue     float64 `json:"totalbuyvalue" example:"0.0"`
	DTotalSellValue    float64 `json:"totalsellvalue" example:"0.0"`
	DCfBuyAvgPrice     float64 `json:"buyavgprice" example:"0.0"`
	DCfSellAvgPrice    float64 `json:"sellavgprice" example:"0.0"`
	DTotalBuyAvgPrice  float64 `json:"totalbuyavgprice" example:"0.0"`
	DTotalSellAvgPrice float64 `json:"totalsellavgprice" example:"0.0"`
	DNetPrice          float64 `json:"netprice" example:"0.0"`
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