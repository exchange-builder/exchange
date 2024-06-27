package model

import "time"

type OrderReq struct {
	// json tag to de-serialize json body
	OrderType string  `json:"orderType"`
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	Quantity  float64 `json:"quantity"`
	Did       float64 `json:"bid"`
	BuyOrSell string  `json:"buyOrSell"`
}

type OrderDto struct {
	// json tag to de-serialize json body
	Id         int64     `db:"id" json:"id"`
	Uid        int64     `db:"uid" json:"uid"`
	OrderType  string    `db:"order_type" json:"orderType"`
	OrderCode  string    `db:"order_code" json:"orderCode"`
	Symbol     string    `db:"symbol" json:"symbol"`
	Amount     float64   `db:"amount" json:"amount"`
	Quantity   float64   `db:"quantity" json:"quantity"`
	Did        float64   `db:"bid" json:"bid"`
	BuyOrSell  int8      `db:"buy_or_sell" json:"buyOrSell"`
	CreateTime time.Time `db:"create_time" json:"createTime"`
	UpdateTime time.Time `db:"update_time" json:"updateTime"`
}
