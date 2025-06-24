package main

import "fmt"

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
  //for loops
  fmt.Println("--------------------")
  for i := 0; i < 5; i++ {
		fmt.Println(i)
  }
  i := 0
  for i < 5 {
    fmt.Println(i)
    i++
  }
  
  for {
    fmt.Println("Sonsuz Döngü")
    break
  }
}