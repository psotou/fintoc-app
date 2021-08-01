package main

import (
	"fintoc-app/pkg/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

const ReqUrl string = "https://api.fintoc.com/v1/accounts/%s/movements?link_token=%s&per_page=%s"

type Movement struct {
	Id              string  `json:"id"`
	Object          string  `json:"object"`
	Amount          float64 `json:"amount"`
	PostDate        string  `json:"post_date"`
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	Currency        string  `json:"currency"`
	ReferenceId     string  `json:"reference_id"`
	Type            string  `json:"type"`
	Pending         bool    `json:"pending"`
	// RecipientAccount string `json:recipient_account"`
	SenderAccount Sender `json:"sender_account"`
	Comment       string `json:"comment"`
}

type Sender struct {
	HolderId    string `json:"holder_id"`
	HolderName  string `json:"holder_name"`
	Number      string `json:"number"`
	Institution Inst   `json:"institution"`
}

type Inst struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	var (
		perPage = flag.String("p", "10", "number of elements per page")
		account = flag.String("c", "cc", "account type: cc or tc")
	)
	flag.Parse()

	var (
		movements []Movement
		accountId string
		amountUSD float64
		date      string
	)
	if *account == "cc" {
		accountId = os.Getenv("BCH_ACCOUNT_ID")
	} else {
		accountId = os.Getenv("BCH_TC_ACCOUNT_ID")
	}
	url := fmt.Sprintf(ReqUrl, accountId, os.Getenv("FINTOC_TOKEN"), *perPage)
	resData := utils.GetReq(url)
	utils.ParseJson(resData, &movements)

	for i := 0; i < len(movements); i++ {
		if *account == "cc" {
			amountUSD = movements[i].Amount
			date = strings.Replace(movements[i].TransactionDate[:19], "T", " ", -1)
		} else {
			amountUSD = movements[i].Amount / 100
			date = strings.Replace(movements[i].PostDate[:19], "T", " ", -1)
		}

		fmt.Println("-------------------------------------------------")
		fmt.Printf("%49s\n", movements[i].Description)
		fmt.Printf("Fecha %43s\n", date)
		fmt.Printf("Monto %43v\n", amountUSD)
		fmt.Printf("Moneda %42s\n", movements[i].Currency)
	}
}
