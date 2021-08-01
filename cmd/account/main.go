package main

import (
	"fintoc-app/pkg/utils"
	"fmt"
	"os"
)

const ReqUrl string = "https://api.fintoc.com/v1/accounts?link_token=%s"

type Account struct {
	Id             string  `json:"id"`
	Type           string  `json:"type"`
	Number         string  `json:"number"`
	Name           string  `json:"name"`
	OfficialName   string  `json:"official_name"`
	AccountBalance Balance `json:"balance"`
	HolderId       string  `json:"holder_id"`
	HolderName     string  `json:"holder_name"`
	Currency       string  `json:"currency"`
	RefreshedAt    string  `json:"refreshed_at"`
	Object         string  `json:"object"`
}

type Balance struct {
	Available float64 `json:"available"`
	Current   float64 `json:"current"`
	Limit     float64 `json:"limit"`
}

func main() {
	var accounts []Account
	resData := utils.GetReq(fmt.Sprintf(ReqUrl, os.Getenv("FINTOC_TOKEN")))
	utils.ParseJson(resData, &accounts)

	usadoUSD := (accounts[5].AccountBalance.Limit - accounts[5].AccountBalance.Available) / 100

	// Cuenta corriente
	fmt.Printf("Cuenta %22s\n", accounts[2].Name)
	fmt.Printf("Disponible %18v\n", accounts[2].AccountBalance.Available)
	fmt.Printf("Moneda %22s\n\n", accounts[2].Currency)
	// Tarjeta de crédito internacional
	fmt.Printf("Cuenta %22s\n", accounts[5].Name)
	fmt.Printf("Usado %23v\n", usadoUSD)
	fmt.Printf("Moneda %22s\n", accounts[5].Currency)
}
