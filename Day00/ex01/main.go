package main

import "fmt"

func someFunction() (int, error) {
	return 42, nil
}

func main() {

	// Zero values
	var nbr int
	fmt.Println(nbr)
	var str string
	fmt.Println(str)
	var boolean bool
	fmt.Println(boolean)
	var float float64
	fmt.Println(float)

	// Multiple variable assignment
	var a, b, c int = 1, 5, 7
	fmt.Printf("first %d, second %d, third %d\n", a, b, c)

	// Blank identifier
	// Only need the error, discard the result
	_, err := someFunction()
	fmt.Println("err:", err)

	// Only need the value, ignore the error
	value, _ := someFunction()
	fmt.Println("value:", value)

	// Only need the index in a range loop
	slice := []int{10, 20, 30}
	for i, _ := range slice {
		fmt.Println("index:", i)
	}

	// Only need the value in a range loop
	for _, value := range slice {
		fmt.Println("value:", value)
	}

	// Go'da alt çizgi _ ile gösterilen boş tanımlayıcı,
	// ihtiyacınız olmayan değerleri atmak için kullanılan özel bir tanımlayıcıdır.
}
