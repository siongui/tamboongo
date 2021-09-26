// Package tamboongo reads donation data and makes donations via Omise Token API
// and Charge API.
package tamboongo

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/omise/omise-go"
)

// MakeDonations makes donations from row data and print summary at the end.
func MakeDonations(records []CsvRecord, verbose bool) (err error) {
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
			if verbose {
				log.Println(err)
			}
			failureDonatedAmount += record.AmountSubunits
			failureCsvRecords = append(failureCsvRecords, record)
			continue
		}
		if verbose {
			log.Printf("created card: %#v\n", card)
		}

		// make donation by creating a charge for the token (credit card)
		charge, err := CreateCharge(client, record.AmountSubunits, "thb", card.Base.ID)
		if err != nil {
			if verbose {
				log.Println(err)
			}
			failureDonatedAmount += record.AmountSubunits
			failureCsvRecords = append(failureCsvRecords, record)
			continue
		}
		if verbose {
			log.Printf("created charge: %#v\n", charge)
		}

		successCsvRecords = append(successCsvRecords, record)
		successDonatedAmount += charge.Amount
	}

	fmt.Println("done.")

	printSummary(successDonatedAmount, failureDonatedAmount, successCsvRecords, len(records))
	return
}

func chargeRoutine(client *omise.Client, record CsvRecord, cs, cf chan CsvRecord, verbose bool) {
	// create a token (credit card)
	card, err := CreateToken(client, record)
	if err != nil {
		if verbose {
			log.Println(err)
		}
		cf <- record
		return
	}
	if verbose {
		log.Printf("created card: %#v\n", card)
	}

	// make donation by creating a charge for the token (credit card)
	charge, err := CreateCharge(client, record.AmountSubunits, "thb", card.Base.ID)
	if err != nil {
		if verbose {
			log.Println(err)
		}
		cf <- record
		return
	}
	if verbose {
		log.Printf("created charge: %#v\n", charge)
	}

	cs <- record
}

var apiAccessInterval = 200 * time.Millisecond

// SetApiAccessInterval sets the interval (millisecond) between API access.
func SetApiAccessInterval(ms int64) {
	apiAccessInterval = time.Duration(ms) * time.Millisecond
}

// MakeConcurrentDonations is the same as MakeDonations, except goroutine is
// used for access Omise API concurrently.
func MakeConcurrentDonations(records []CsvRecord, verbose bool) (err error) {
	fmt.Println("performing donations...")

	client, err := NewOmiseClient()
	if err != nil {
		return
	}

	var failureCsvRecords []CsvRecord
	var successCsvRecords []CsvRecord
	var successDonatedAmount int64 = 0
	var failureDonatedAmount int64 = 0

	cSuccess := make(chan CsvRecord)
	cFailure := make(chan CsvRecord)

	for _, record := range records {
		go chargeRoutine(client, record, cSuccess, cFailure, verbose)
		time.Sleep(apiAccessInterval)
	}

	for i := 0; i < len(records); i++ {
		select {
		case rs := <-cSuccess:
			successDonatedAmount += rs.AmountSubunits
			successCsvRecords = append(successCsvRecords, rs)

		case rf := <-cFailure:
			failureDonatedAmount += rf.AmountSubunits
			failureCsvRecords = append(failureCsvRecords, rf)
		}
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
	if len(sRecords) != 0 {
		fmt.Printf("average per person:\t THB\t %d\n", sAmount/int64(len(sRecords)))
	}

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
