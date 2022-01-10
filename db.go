package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Balances struct {
	name   string
	amount int64
}

type Genesis struct {
	genesisTime time.Time
	chainId     string
	balances    []Balances
}

func writeGenesis() {
	genesis := Genesis{
		genesisTime: time.Now(),
		chainId:     "the-blockchain-bar-ledger",
		balances: []Balances{
			{
				name:   "andrej",
				amount: 1000000,
			},
		},
	}
	content, err := json.Marshal(genesis)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("database/genesis.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
