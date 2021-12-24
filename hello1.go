package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fun1()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
}
func fun1() {
	fmt.Print("działanie funkcji")
}
