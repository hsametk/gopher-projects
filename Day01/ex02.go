package main

import "fmt"
func main() {
    i := 0

start: // Label tanımı
    fmt.Printf("i = %d\n", i)
    i++
    
    if i < 5 {
        goto start // Label'a atla
    }
    
    fmt.Println("Döngü bitti")
}