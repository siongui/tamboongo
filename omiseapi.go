package tamboongo

import (
	"errors"
	"os"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func NewOmiseClient() (client *omise.Client, err error) {
	pkey, okpkey := os.LookupEnv("OmisePublicKey")
	if !okpkey {
		err = errors.New("OmisePublicKey environment variable not set")
		return
	}

	skey, okskey := os.LookupEnv("OmiseSecretKey")
	if !okskey {
		err = errors.New("OmiseSecretKey environment variable not set")
		return
	}

	client, err = omise.NewClient(pkey, skey)
	return
}

func CreateToken(client *omise.Client, record CsvRecord) (card *omise.Card, err error) {
	card, createToken := &omise.Card{}, &operations.CreateToken{
		Name:            record.Name,
		Number:          record.CCNumber,
		ExpirationMonth: record.ExpMonth,
		ExpirationYear:  record.ExpYear,
		SecurityCode:    record.CVV,
	}

	err = client.Do(card, createToken)
	return
}

func CreateCharge(client *omise.Client, amount int64, currency, tokenid string) (charge *omise.Charge, err error) {
	charge, create := &omise.Charge{}, &operations.CreateCharge{
		Amount:   amount,
		Currency: currency,
		Card:     tokenid,
	}

	err = client.Do(charge, create)
	return
}
