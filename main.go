package main

import (
	"fmt"
	"log"

	"github.com/mikunup/test/fizzbuzz"
)

func main() {
	r, err := fizzbuzz.FizzBuzz(15)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf(r)
}
