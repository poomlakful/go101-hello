package main

import "fmt"

func main() {
	renderFibonacci(10)
}

func renderFibonacci(num int) {
	add, print := fibonacci()
	for i := 0; i < num; i++ {
		print()
		add()
	}
}

func fibonacci() (func(), func()) {
	a, b := 0, 1
	return func() {
			a, b = b, a+b
		},
		func() {
			fmt.Println(a)
		}
}
