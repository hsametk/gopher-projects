package main

import "fmt"

func multiple(n int)  {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", n, i, (n* i))
	}
}
func main()  {
	var i int
	fmt.Println("input number")
	fmt.Scan(&i)
	multiple(i)
}