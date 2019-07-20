package fizzbuzz_test

/**
*
*  package test name should be end with ${package_name}_test.go
*
**/

import (
	"hello/fizzbuzz"
	"testing"
)

/**
*
*	1 = "1"
*	2 = "2"
*	3 = "Fizz"
*	4 = "4"
*	5 = "5"
*	6 = "Buzz"
*
**/
func TestFizzBuzzNumber(t *testing.T) {
	t.Run("hello test 1", func(t *testing.T) {
		testFizzBuzz("1", 1, t)
	})

	t.Run("hello test 2", func(t *testing.T) {
		testFizzBuzz("2", 2, t)
	})
}

func testFizzBuzz(want string, given int, t *testing.T) {
	get := fizzbuzz.FizzBuzz(given)

	if want != get {
		t.Errorf("it should be %s but get %s", want, get)
	}
}
