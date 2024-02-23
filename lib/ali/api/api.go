package api

import (
	"github.com/go-resty/resty/v2"
)

type ExchangeSingleRes struct {
	Result struct {
		List map[string]Exchange `json:"list"`
	} `json:"result"`
}

type Exchange struct {
	Name       string `json:"name"`
	Rate       string `json:"rate"`
	UpdateTime string `json:"updatetime"`
}

// ExchangeSingle 单个货币汇率查询
func ExchangeSingle(appCode, currency string) (ExchangeSingleRes, error) {
	var result ExchangeSingleRes
	_, err := resty.New().R().
		SetHeader("Authorization", "APPCODE "+appCode).
		SetQueryParam("currency", currency).
		SetResult(&result).
		Get("https://jisuhuilv.market.alicloudapi.com/exchange/single")
	return result, err
}
