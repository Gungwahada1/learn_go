package functions

import "fmt"

var c int

func TestingFunction(a int, b int) {
	fmt.Println("Hallo, ini dari fungsi testing")

	c = a + b

	fmt.Printf("a tambah b sama dengan %d \n", c)
}
