package main

import "fmt"

func main() {
	ninja_2_1()
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
}
func ninja_1_4() {
	type typex int
	var x typex
	x = 42
	fmt.Printf("%T--%v", x, x)
}
func ninja_1_5() {
	type typex int
	var x typex
	var y int
	x = 42
	y = int(x)
	fmt.Printf("%T--%v", x, x)
	fmt.Printf("\n%T--%v", y, y)
}
func ninja_2_1() {
	fmt.Println("ssas", 87)

}
