package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	resData := GetReq(fmt.Sprintf(ReqUrl, os.Getenv("FINTOC_TOKEN")))
	ParseJson(resData, &accounts)

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

func GetReq(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("Authorization", os.Getenv("FINTOC_LIVE_SECRET"))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	return body
}

func ParseJson(data []byte, bankData interface{}) interface{} {
	err := json.Unmarshal(data, &bankData)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bankData
}
