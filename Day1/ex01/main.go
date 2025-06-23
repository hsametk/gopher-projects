package main

import "fmt"

func main()  {
	
	//Zero_values
	var nbr int
	fmt.Println(nbr)
	var str string
	fmt.Println(str)
	var boolean bool
	fmt.Println(boolean)
	var float float64
	fmt.Println(float)

	//Multiple variable assignment
	var a,b,c int  = 1, 5, 7
	fmt.Printf("first %d, second %d, third %d", a,b,c)
	//Blank identifier
	// Only need the error, discard the result
	_, err := someFunction()

	// Only need the value, ignore the error
	value, _ := someFunction()

	// Only need the index in a range loop
	for i, _ := range slice {
	    // use i, ignore the value
	}

	// Only need the value in a range loop
	for _, value := range slice {
	    // use value, ignore the index
	}
	//Go'da alt çizgi _ ile gösterilen boş tanımlayıcı, 
	// ihtiyacınız olmayan değerleri atmak için kullanılan özel bir tanımlayıcıdır.
}