package main

import (
	"fintoc-app/pkg/utils"
	"fmt"
	"os"
)

const ReqUrl string = "https://api.fintoc.com/v1/accounts/%s/movements?link_token=%s"

type Movement struct {
	Id              string `json:"id"`
	Object          string `json:"object"`
	Amount          int    `json:"amount"`
	PostDate        string `json:"post_date"`
	Description     string `json:"description"`
	TransactionDate string `json:"transaction_date"`
	Currency        string `json:"currency"`
	ReferenceId     string `json:"reference_id"`
	Type            string `json:"type"`
	Pending         bool   `json:"pending"`
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
	var movements []Movement
	url := fmt.Sprintf(ReqUrl, os.Getenv("BCH_ACCOUNT_ID"), os.Getenv("FINTOC_TOKEN"))
	resData := utils.GetReq(url)
	utils.ParseJson(resData, &movements)

	for i := 0; i < len(movements); i++ {
		fmt.Printf("Fecha %43s\n", movements[i].TransactionDate)
		fmt.Printf("Descripción %37s\n", movements[i].Description)
		fmt.Printf("Monto %43v\n", movements[i].Amount)
		fmt.Printf("Moneda %42s\n", movements[i].Currency)
		fmt.Printf("Pendiente %39v\n", movements[i].Pending)
	}

}
