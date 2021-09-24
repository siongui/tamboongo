package tamboongo

import (
	"strings"
	"testing"
)

func TestDecryptRot(t *testing.T) {
	b, err := DecryptRot("fng.1000.csv.rot128")
	if err != nil {
		t.Error(err)
		return
	}

	if !strings.HasPrefix(string(b), "Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear") {
		t.Error(string(b))
		return
	}
}

func TestReadCsv(t *testing.T) {
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

	for _, record := range records {
		if len(record) != 6 {
			t.Error(record)
			return
		}
	}
}
