package tamboongo

import (
	"errors"
	"log"
	"os"
	"time"

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

func CreateToken(client *omise.Client, name, number, cvv string, month time.Month, year int) (err error) {
	card, createToken := &omise.Card{}, &operations.CreateToken{
		Name:            name,
		Number:          number,
		ExpirationMonth: month,
		ExpirationYear:  year,
		SecurityCode:    cvv,
	}

	if err = client.Do(card, createToken); err != nil {
		return
	}

	log.Printf("created card: %#v\n", card)

	return
}
