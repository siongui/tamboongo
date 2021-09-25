package main

import (
	"flag"
	"fmt"

	"github.com/siongui/tamboongo"
)

var rot = flag.String("rot", "", "A ROT-128 encrypted CSV file")

func main() {
	flag.Parse()

	b, err := tamboongo.DecryptRot(*rot)
	if err != nil {
		fmt.Println(err)
		return
	}

	records, err := tamboongo.ReadCsv(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tamboongo.MakeDonations(records)
	if err != nil {
		fmt.Println(err)
		return
	}
}
