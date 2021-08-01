package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

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
