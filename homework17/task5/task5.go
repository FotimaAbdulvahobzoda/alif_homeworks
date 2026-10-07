package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Login    string  `json:"login"`
	Password string  `json:"-"`
	Balance  float64 `json:"balance"`
}

func main() {
	acc := Account{
		Login:    "Rebbeca_Wood",
		Password: "grek7UUop",
		Balance:  56.780,
	}

	data, _ := json.MarshalIndent(acc, "", " ")
	fmt.Println(string(data))
}
