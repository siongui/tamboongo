package main

import (
	"flag"
	"fmt"

	"github.com/siongui/tamboongo"
)

var rot = flag.String("rot", "", "A ROT-128 encrypted CSV file")
var verbose = flag.Bool("verbose", false, "print more messages while performing actions")

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

	err = tamboongo.MakeDonations(records, *verbose)
	if err != nil {
		fmt.Println(err)
		return
	}
}
