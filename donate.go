package tamboongo

import (
	"log"
)

func MakeDonations(records []CsvRecord) (err error) {
	client, err := NewOmiseClient()
	if err != nil {
		return
	}

	for _, record := range records {
		card, err := CreateToken(client, record)
		if err != nil {
			log.Println(err)
			continue
		}

		//log.Printf("created card: %#v\n", card)
		//log.Println(card)
		charge, err := CreateCharge(client, record.AmountSubunits, "thb", card.Base.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("created charge: %#v\n", charge)
	}

	return
}
