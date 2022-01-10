package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Balances struct {
	Name   string
	Amount int64
}

type Genesis struct {
	GenesisTime time.Time
	ChainId     string
	Balances    []Balances
}

func writeGenesis() {
	gen := Genesis{
		GenesisTime: time.Now(),
		ChainId:     "the-blockchain-bar-ledger",
		Balances: []Balances{
			{
				Name:   "andrej",
				Amount: 1000000,
			},
		},
	}

	file, _ := json.MarshalIndent(gen, "", " ")

	_ = ioutil.WriteFile("database/genesis.json", file, 0644)
}

func writeBalances() {
	balances := []Balances{
		{
			Name:   "andrej",
			Amount: 1000700,
		},
	}
	file, _ := json.MarshalIndent(balances, "", " ")

	_ = ioutil.WriteFile("database/state.json", file, 0644)
}
