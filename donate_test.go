package tamboongo

import (
	"testing"
)

func ExampleMake20Donations(t *testing.T) {
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

	err = MakeDonations(records[:20])
	if err != nil {
		t.Error(err)
		return
	}
}
