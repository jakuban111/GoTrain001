package main

import "fmt"

func main() {
	ninja_1_3()
}
func ninja_1_1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)
}
func ninja_1_2() {
	var x int
	var y string
	var z bool
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
func ninja_1_3() {
	x := 42
	y := "James Bond"
	z := true
	var s string
	s = fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Printf(s)
	//s := fmt.Sprintf(x, y, z)
	//fmt.Println(s)
}
