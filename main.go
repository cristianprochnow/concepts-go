package main

import "fmt"

func main() {
	// Values
	var text string = "This is a string"
	var numbers int32 = 1 + 1
	var decimal float32 = 7.2
	var booleans bool = true

	fmt.Println(fmt.Sprintf("string: %s", text))
	fmt.Println(fmt.Sprintf("int: %d", numbers))
	fmt.Println(fmt.Sprintf("float: %f", decimal))
	fmt.Println(fmt.Sprintf("bool: %t", booleans))

	// Variables
	var a = "Start"
	var b, c = 1, 2
	var d = true
	var e int

	fmt.Println(a, b, c, d, e)
}
