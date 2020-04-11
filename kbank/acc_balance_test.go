package kbank_test

import (
	"github.com/tanopwan/thaibankgo/kbank"
	"testing"
)

func TestApp_GetBalance(t *testing.T) {
	config := kbank.AppConfig{
		PartnerID:     "{{YOUR TXN ID}}",
		PartnerSecret: "{{YOUR PARTNER ID}}",
		MerchantID:    "{{YOUR PARTNER SECRET}}",
	}
	app := kbank.NewKBankApp(config)
	response, err := app.GetBalance("1111111111")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("result: %v\n", response)
}
