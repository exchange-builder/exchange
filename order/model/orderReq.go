package model

type OrderReq struct {
	// json tag to de-serialize json body
	OrderType string  `json:"orderType"`
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	Quantity  float64 `json:"quantity"`
	Did       float64 `json:"bid"`
	BuyOrSell string  `json:"buyOrSell"`
}
