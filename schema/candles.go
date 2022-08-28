package schema

import (
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	Date   time.Time       `json:"date"`
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume uint64          `json:"volume"`
}

type CandlesResponse struct {
	SymbolID string    `json:"symbolId"`
	Type     string    `json:"type"`
	Exchange string    `json:"exchange"`
	Market   string    `json:"market"`
	Candles  []*Candle `json:"candles"`
}
