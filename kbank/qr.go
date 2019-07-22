package kbank

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type KBankApp struct {
	config KBankAppConfig
}

type KBankAppConfig struct {
	PartnerID     string
	PartnerSecret string
	MerchantID    string
}

func NewKBankApp(config KBankAppConfig) *KBankApp {
	return &KBankApp{
		config: config,
	}
}

func generateID() string {
	b := make([]byte, 7)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

type Money float64

func (m Money) MarshalJSON() ([]byte, error) {
	// There are probably better ways to do it. It is just an example
	return []byte(fmt.Sprintf("%.2f", m)), nil
}

type Payload struct {
	PartnerTxnUID   string  `json:"partnerTxnUid"`
	PartnerID       string  `json:"partnerId"`
	PartnerSecret   string  `json:"partnerSecret"`
	RequestDt       string  `json:"requestDt"`
	MerchantID      string  `json:"merchantId"`
	TerminalID      string  `json:"terminalId"`
	QrType          string  `json:"qrType"`
	TxnAmount       Money   `json:"txnAmount"`
	TxnCurrencyCode string  `json:"txnCurrencyCode"`
	Reference1      string  `json:"reference1"`
	Reference2      *string `json:"reference2"`
	Reference3      *string `json:"reference3"`
	Reference4      *string `json:"reference4"`
	Metadata        string  `json:"metadata"`
}

type Response struct {
	PartnerTxnUID string `json:"partnerTxnUid"`
	PartnerID     string `json:"partnerId"`
	StatusCode    string `json:"statusCode"`
	ErrorCode     string `json:"errorCode"`
	ErrorDesc     string `json:"errorDesc"`
	AccountName   string `json:"accountName"`
	QRCode        string `json:"qrCode"`
}

func (a *KBankApp) GenerateQR(amount Money, reference1, metadata string) (Response, error) {
	data := Payload{
		// fill struct
		PartnerTxnUID:   generateID(),
		PartnerID:       a.config.PartnerID,
		PartnerSecret:   a.config.PartnerSecret,
		RequestDt:       time.Now().Format(time.RFC3339),
		MerchantID:      a.config.MerchantID,
		TerminalID:      "term1",
		QrType:          "3",
		TxnAmount:       amount,
		TxnCurrencyCode: "THB",
		Reference1:      reference1,
		Metadata:        metadata,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to prepare request body with reason: %s\n", err.Error())
		return Response{}, err
	}
	body := bytes.NewReader(payloadBytes)
	log.Printf("body: %s", string(payloadBytes))

	req, err := http.NewRequest("POST", "https://APIPORTAL.kasikornbank.com:12002/pos/qr_request", body)
	if err != nil {
		log.Printf("failed to preapre new request KBank Api with reason: %s\n", err.Error())
		return Response{}, err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to call KBank Api with reason: %s\n", err.Error())
		return Response{}, err
	}
	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response from KBank Api with reason: %s\n", err.Error())
		return Response{}, err
	}

	fmt.Printf(string(bb))
	var response Response
	err = json.Unmarshal(bb, &response)
	if err != nil {
		log.Printf("failed to unmarshal response from KBank Api with reason: %s\n", err.Error())
		return Response{}, err
	}

	return response, nil
}
