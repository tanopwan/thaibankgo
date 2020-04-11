package kbank

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var balanceURL = bankURL + "/deposit/sight/balance/account"
var transactionURL = bankURL + "/deposit/sight/transactions/account"

type BalanceResponse struct {
	AccID        string  `json:"accId"`
	AvailBalance float64 `json:"availBalance"`
	AcctBalance  float64 `json:"acctBalance"`
	AcctStatus   string  `json:"acctStatus"`
}

type TransactionResponse struct {
	TotalItems int `json:"totalItems"`
	Items      []struct {
		ToAccountNo        string  `json:"toAccountNo"`
		ToAccountName      string  `json:"toAccountName"`
		ToAccountNameEN    string  `json:"toAccountNameEN"`
		ToAccountNameTH    string  `json:"toAccountNameTH"`
		ChannelDetail      string  `json:"channelDetail"`
		MerchantCode       string  `json:"merchantCode"`
		FromAccountID      string  `json:"fromAccountId"`
		FromAccountNameTH  string  `json:"fromAccountNameTH"`
		FromAccountNameEN  string  `json:"fromAccountNameEN"`
		OutstandingBalance float64 `json:"outstandingBalance"`
		FromBankCode       string  `json:"fromBankCode"`
		TxnTime            string  `json:"txnTime"`
		TxnAmount          float64 `json:"txnAmount"`
		FeeAmount          int     `json:"feeAmount"`
		ServiceBranchNo    string  `json:"serviceBranchNo"`
		TxnDescEN          string  `json:"txnDescEN"`
		EffectiveDate      string  `json:"effectiveDate"`
		ChannelCode        string  `json:"channelCode"`
		TxnDate            string  `json:"txnDate"`
		TxnDesc            string  `json:"txnDesc"`
		ChequeNo           string  `json:"chequeNo"`
		ProxyID            string  `json:"proxyId"`
		ProxyTypeCode      string  `json:"proxyTypeCode"`
		ToBankCode         string  `json:"toBankCode"`
		TellerID           string  `json:"tellerId"`
		DebitCreditFlag    string  `json:"debitCreditFlag"`
	} `json:"items"`
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

func (a *App) GetTransaction(account string) (TransactionResponse, error) {
	u, err := url.Parse(transactionURL)
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
		return TransactionResponse{}, err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Partner-Id", a.config.PartnerID)
	req.Header.Set("Partner-Secret", a.config.PartnerSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to call KBank Api with reason: %s\n", err.Error())
		return TransactionResponse{}, err
	}
	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response from KBank Api with reason: %s\n", err.Error())
		return TransactionResponse{}, err
	}

	log.Println(string(bb))
	var response TransactionResponse
	err = json.Unmarshal(bb, &response)
	if err != nil {
		log.Printf("failed to unmarshal response from KBank Api with reason: %s\n", err.Error())
		return TransactionResponse{}, err
	}

	return response, nil
}
