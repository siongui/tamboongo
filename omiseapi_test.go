package tamboongo

import (
	"testing"
)

func ExampleCreateToken(t *testing.T) {
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
		if index == 0 {
			t.Log(record)
			continue
		}
		if index > 5 {
			break
		}

		t.Log(record)

		err = CreateToken(client, record)
		if err != nil {
			t.Error(err)
		}
	}
}
