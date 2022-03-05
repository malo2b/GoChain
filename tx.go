package main

type Account string

func newAccount(value string) Account {
	return Account(value)
}

type Tx struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
}

func newTx(from Account, to Account, value uint, data string) Tx {
	return Tx{from, to, value, data}
}

func (t Tx) IsReward() bool {
	return t.Data == "reward"
}

type Balances struct {
	Name   string
	Amount int64
}
