package main

import (
	"fmt"
)

func main()  {
	//if else

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
	fmt.Printf("------------\n")
	time := 20
	if (time < 18) {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening")
	}
}

// ex00: if_else           - Conditional statements DONE
// ex01: switch_case       - Switch statements (with and without expression) DONE
// ex02: for_loops         - All three forms of for loops DONE
// ex03: range_iteration   - Range over slices, maps, strings  TODO
// ex04: break_continue    - Loop control statements
// ex05: goto_labels       - Proper use of goto (rare cases)
// ex06: fizzbuzz          - Classic FizzBuzz problem
// ex07: prime_checker     - Check if number is prime
// ex08: multiplication_table - Generate multiplication tables