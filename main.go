package main

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"github.com/sleepypikachu/still-good-food/food"
=======
	"github.com/sleepypikachu/food"
>>>>>>> 27ac30aa7d8c1642c8692e2708197bdfca74f587
	"os"
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

	fmt.Printf("%s\n", b)
}
