package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	r.GET("/balancesList", func(c *gin.Context) {
		c.JSON(200,
			balancesListJson(),
		)
	})

	r.POST("/addTx", func(c *gin.Context) {
		u64, err := strconv.ParseUint(c.PostForm("value"), 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		value := uint(u64)
		txAdd(c.PostForm("from"), c.PostForm("to"), value)
		c.String(200, "La transaction a bien été enregistrée")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//balancesList()
}

func txAdd(from string, to string, value uint) {

	fromAcc := newAccount(from)
	toAcc := newAccount(to)

	tx := newTx(fromAcc, toAcc, value, "")

	state, err := NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// defer means, at the end of this function execution,
	// execute the following statement (close DB file with all TXs)
	defer state.Close()

	// Add the TX to an in-memory array (pool)
	err = state.add(tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Flush the mempool TXs to disk
	err = state.persist()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("TX successfully added to the ledger.")
}

func balancesList() {
	state, err := NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	fmt.Println("Accounts balances:")
	fmt.Println("__________________")
	fmt.Println("")
	for account, balance := range state.Balances {
		fmt.Println(fmt.Sprintf("%s: %d", account, balance))
	}
}

func balancesListJson() []*Balance {
	state, err := NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	var result []*Balance
	for account, balance := range state.Balances {
		result = append(result, &Balance{Account: string(account), Balancee: uint(balance)})

		fmt.Println(fmt.Sprintf("%s: %d", account, balance))
	}

	return result
}

type Balance struct {
	Account  string `json:"account"`
	Balancee uint   `json:"balance"`
}
