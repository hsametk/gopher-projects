package main

import "fmt"

const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
func main() {
	var firstname = "samet"
	var lastName string = "kotu"
	number := 123
	var tf bool = true;
	fmt.Println("Hello world!")
	fmt.Println(firstname, lastName)
	fmt.Println(number);
	fmt.Println(tf);
	//------------------------
	const age int = 23;
	fmt.Println(age);
	//-------------
	fmt.Println(Sunday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	//-----------------------
	var i int = 42
	var f float64 = float64(i)
	u := uint(f)
	fmt.Println("\n\n\n")
	fmt.Println("int:", i)
	fmt.Println("float:", f)
	fmt.Println("uint:", u)
}

// ex00: hello_world         - Print "Hello, World!" DONE
// ex01: variables          - Declare and use different variable types DONE
// ex02: constants          - Work with constants and iota DONE
// ex03: basic_types        - Explore all basic types
// ex04: type_conversion    - Convert between types safely DONE
// ex05: zero_values        - Understand zero values DONE
// ex06: multiple_assignment - Multiple variable assignment DONE
// ex07: blank_identifier   - Use blank identifier properly DONE
