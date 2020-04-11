package kbank_test

import (
	"testing"

	"github.com/tanopwan/thaibankgo/kbank"
)

func TestMarshalMoney(t *testing.T) {
	money := (kbank.Money)(100.50)
	result, err := money.MarshalJSON()
	if err != nil {
		t.Errorf("err when marshal json with reason: %s", err.Error())
	}

	if string(result) != "100.50" {
		t.Errorf("expect 2 decimal places got %s", string(result))
	}
	t.Logf("result: %s\n", string(result))
}

func TestApp_GenerateQR(t *testing.T) {
	config := kbank.AppConfig{
		PartnerID:     "{{YOUR TXN ID}}",
		PartnerSecret: "{{YOUR PARTNER ID}}",
		MerchantID:    "{{YOUR PARTNER SECRET}}",
	}
	app := kbank.NewKBankApp(config)
	response, err := app.GenerateQR(100.50, "INV001", "ถุงผ้า 80.50, ดินสอ 20.00")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("result: %v\n", response)
}
