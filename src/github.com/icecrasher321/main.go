package main

import (
	"encoding/json"
	"fmt"
	"os"
 "github.com/icecrasher321/food"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <url>\n", os.Args[0])
		return
	}

	r, err := food.Scrape(os.Args[1])
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)
}
