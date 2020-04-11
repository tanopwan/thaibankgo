package kbank

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var balanceURL = bankURL + "/deposit/sight/balance/account"

type BalanceResponse struct {
	AccID        string  `json:"accId"`
	AvailBalance float64 `json:"availBalance"`
	AcctBalance  float64 `json:"acctBalance"`
	AcctStatus   string  `json:"acctStatus"`
}

func (a *App) GetBalance(account string) (BalanceResponse, error) {
	u, err := url.Parse(balanceURL)
	if err != nil {
		log.Fatal(err)
	}
	rel, err := u.Parse(account)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", rel.String(), nil)
	if err != nil {
		log.Printf("failed to preapre new request KBank Api with reason: %s\n", err.Error())
		return BalanceResponse{}, err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Partner-Id", a.config.PartnerID)
	req.Header.Set("Partner-Secret", a.config.PartnerSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to call KBank Api with reason: %s\n", err.Error())
		return BalanceResponse{}, err
	}
	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response from KBank Api with reason: %s\n", err.Error())
		return BalanceResponse{}, err
	}

	log.Println(string(bb))
	var response BalanceResponse
	err = json.Unmarshal(bb, &response)
	if err != nil {
		log.Printf("failed to unmarshal response from KBank Api with reason: %s\n", err.Error())
		return BalanceResponse{}, err
	}

	return response, nil
}
