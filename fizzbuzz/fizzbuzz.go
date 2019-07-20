package fizzbuzz

import "fmt"

// FizzBuzz is Converted to FizzBuzzWord
func FizzBuzz(n int) (string, error) {
	if n < 1 || n > 100 {
		return "", fmt.Errorf("invalid number:%v", n)
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
