package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/image/font/sfnt"
)

func main() {
	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	f, err := sfnt.Parse(src)
	if err != nil {
		panic(err)
	}

	names := []sfnt.NameID{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		13,
		14,
		16,
		17,
		18,
		19,
		20,
		21,
		22,
		23,
		24,
		25,
	}

	for _, id := range names {
		value, err := f.Name(nil, id)
		if err != nil {
			continue
		}

		fmt.Printf("%d: %s\n", id, value)
	}
}
