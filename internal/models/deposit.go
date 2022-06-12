package models

type Deposit struct {
	WalletId string  `json:"walletId"`
	Amount   float64 `json:"amount"`
}
