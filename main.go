package main

import (
	"fmt"
)

func main() {
	r, err := fizzBuzz(12)

	if err != nil {
		fmt.Printf("test")
	}
	s := string(r)
	fmt.Printf(s)
}

func fizzBuzz(n int) (string, error) {
	if n < 1 || n > 100 {
		return "", fmt.Errorf("invalid number: $v", n)
	}

	switch {
	case n%15 == 0:
		return "FizzBuzz", nil
	case n%3 == 0:
		return "Fizz", nil
	case n%5 == 0:
		return "Buzz", nil
	default:
		return fmt.Sprint(n), nil
	}
}
