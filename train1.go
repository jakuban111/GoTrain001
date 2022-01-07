package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func alfa() {
	test003()
}
func test001() {

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

}
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func test002(i int) int {
	i = i + 2
	fmt.Println("test002: ", i)
	return i
}
func test003() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
func incrementor() func() int {
	var x int
	fmt.Println("start")
	return func() int {
		x++
		return x
	}
}
