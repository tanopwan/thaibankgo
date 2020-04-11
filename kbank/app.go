package kbank

import (
	"crypto/rand"
	"encoding/hex"
)

var bankURL = "https://APIPORTAL.kasikornbank.com:12002"

type App struct {
	config AppConfig
}

type AppConfig struct {
	PartnerID     string
	PartnerSecret string
	MerchantID    string
}

func NewKBankApp(config AppConfig) *App {
	return &App{
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

