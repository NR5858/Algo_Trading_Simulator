package models

type Order struct {
	Symbol    string `json:"symbol"`
	Action    string `json:"action"`
	Quantity  int    `json:"quantity"`
	PriceType string `json:"priceType"`
	Duration  string `json:"duration"`
}
