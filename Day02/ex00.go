package main

import "fmt"
//basic functions
func  sum_of_two_number(a, b int) int{
	return a + b
}

//multiple return values
func sumAndDiff(a,b int) (int, int)  {
	return a + b, a - b
}
//Named return values
func reactangle(length, width float64) (area, perimeter float64)  {
	area = length * width
	perimeter = 2 * (length * width)
	return area, perimeter	
}
//Variadic funcitons
 func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
func main()  {
	fmt.Println(sum_of_two_number(5,4))
	fmt.Println(sumAndDiff(5, 4))
	fmt.Println(reactangle(5,5))
	fmt.Println(sum(5, 5, 10, 50, 11))
}


// ex00: basic_functions    - Function declaration and calling
// ex01: multiple_returns   - Functions returning multiple values
// ex02: named_returns      - Named return values
// ex03: variadic_functions - Functions with variable arguments
// ex04: higher_order_funcs - Functions as parameters/return values
// ex05: closures          - Anonymous functions and closures
// ex06: recursion         - Recursive functions (factorial, fibonacci)
// ex07: defer_statement   - Proper use of defer
// ex08: calculator        - Build a calculator with functions
// ex09: function_pointers - Function variables and pointers	

