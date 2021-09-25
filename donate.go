// Package tamboongo reads donation data and makes donations via Omise Token API
// and Charge API.
package tamboongo

import (
	"fmt"
	"sort"
)

// MakeDonations makes donations from row data and print summary at the end.
func MakeDonations(records []CsvRecord) (err error) {
	fmt.Println("performing donations...")

	client, err := NewOmiseClient()
	if err != nil {
		return
	}

	var failureCsvRecords []CsvRecord
	var successCsvRecords []CsvRecord
	var successDonatedAmount int64 = 0
	var failureDonatedAmount int64 = 0

	for _, record := range records {
		// create a token (credit card)
		card, err := CreateToken(client, record)
		if err != nil {
			//log.Println(err)
			failureDonatedAmount += record.AmountSubunits
			failureCsvRecords = append(failureCsvRecords, record)
			continue
		}
		//log.Printf("created card: %#v\n", card)

		// make donation by creating a charge for the token (credit card)
		charge, err := CreateCharge(client, record.AmountSubunits, "thb", card.Base.ID)
		if err != nil {
			//log.Println(err)
			failureDonatedAmount += record.AmountSubunits
			failureCsvRecords = append(failureCsvRecords, record)
			continue
		}
		//log.Printf("created charge: %#v\n", charge)

		successCsvRecords = append(successCsvRecords, record)
		successDonatedAmount += charge.Amount
	}

	fmt.Println("done.")

	printSummary(successDonatedAmount, failureDonatedAmount, successCsvRecords, len(records))
	return
}

func printSummary(sAmount, fAmount int64, sRecords []CsvRecord, totalRecordCount int) {
	fmt.Printf("total received:\t\t THB\t %d\n", sAmount+fAmount)
	fmt.Printf("successful donated:\t THB\t %d\n", sAmount)
	fmt.Printf("faulty donation:\t THB\t %d\n", fAmount)
	fmt.Printf("\n")
	fmt.Printf("average per person:\t THB\t %d\n", sAmount/int64(len(sRecords)))

	sort.Slice(sRecords, func(i, j int) bool {
		return sRecords[i].AmountSubunits > sRecords[j].AmountSubunits
	})

	fmt.Printf("top donors:\n")
	for i, record := range sRecords {
		if i > 2 {
			break
		}

		fmt.Printf("\t%s (THB %d)\n", record.Name, record.AmountSubunits)
	}
}
