package tamboongo

import (
	"strconv"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	b, err := DecryptRot("fng.1000.csv.rot128")
	if err != nil {
		t.Error(err)
		return
	}

	records, err := ReadCsv(b)
	if err != nil {
		t.Error(err)
		return
	}

	client, err := NewOmiseClient()
	if err != nil {
		t.Error(err)
		return
	}

	for index, record := range records {
		if len(record) != 6 {
			t.Error(record)
			return
		}
		if index == 0 {
			t.Log(record)
			continue
		}
		if index > 2 {
			break
		}

		t.Log(record)
		name := record[0]
		number := record[2]
		cvv := record[3]

		valMonth, err := strconv.Atoi(record[4])
		if err != nil {
			t.Error(record, " fail to convert month")
			return
		}
		month := time.Month(valMonth)

		year, err := strconv.Atoi(record[5])
		if err != nil {
			t.Error(record, " fail to convert year")
			return
		}
		err = CreateToken(client, name, number, cvv, month, year)
		if err != nil {
			t.Log(err)
			return
		}
	}
}
