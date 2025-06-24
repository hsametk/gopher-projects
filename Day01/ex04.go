package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    
    // Sadece tek sayıları kontrol et, sqrt(n)'e kadar
    for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}
func main()  {
	numbers := []int{2, 3, 4, 5, 17, 25, 29, 97, 100}

	for _, num := range numbers {
		if isPrime(num) {
			fmt.Printf("%d is prime\n", num)
		} else {
			fmt.Printf("%d is not prime\n", num)
		}
	}
}