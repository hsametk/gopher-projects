package main

import "fmt"

func divide(n int) {
	for i := 0; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	var i int
	fmt.Println("Enter number:")
	fmt.Scan(&i)

	divide(i)
}
