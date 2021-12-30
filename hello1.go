package main

import (
	"fmt"
	"time"
)

func main() {
	ninja_4_1()
}
func test01() {
	x := 7
	fmt.Printf("%d\t\t%b\n", x, x)
	x = x << 1
	fmt.Printf("%d\t\t%b\n", x, x)
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
	x := 42
	fmt.Printf("\n%d\t%b\t%#x", x, x, x)
}
func ninja_2_2() {
	x := 42
	y := x != 41
	fmt.Println(y)
}
func ninja_2_3() {
	const (
		a = iota
		b
		b1
		c = 7
		d = iota
		e
	)
	fmt.Println(a, b, b1, c, d, e)
}
func ninja_2_4() {
	var a int = 42
	fmt.Printf("\n%d\t%b\t%#x", a, a, a)
	a = a << 1
	fmt.Printf("\n%d\t%b\t%#x", a, a, a)
}
func ninja_2_5() {
	x := `"d"uP
	\a`
	fmt.Println(x)
}
func ninja_2_6() {
	const (
		a = 2021 + iota
		b = 2021 + iota
		c = 2021 + iota
		d = 2021 + iota
	)
	fmt.Println(a, b, c, d)
}
func ninja_3_1() {
	for i := 0; i < 10000; i++ {
		fmt.Print(i, "\t")
	}
}
func ninja_3_2() {
	for i := 65; i < 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
func ninja_3_3() {
	for i := 1990; i <= time.Now().Year(); i++ {
		fmt.Println(i)
	}
}
func ninja_3_4() {
	x := 1990
	fmt.Println(x)
	for {
		x++
		fmt.Println(x)
		if x >= time.Now().Year() {
			break
		}
	}
}
func ninja_3_5() {
	for i := 10; i < 100; i++ {
		fmt.Println(i, "_", i%4)
	}
}
func ninja_3_6() {
	x := 7
	if x == 6 {
		fmt.Println("true")
	}
}
func ninja_3_7() {
	x := 8
	if x == 6 {
		fmt.Println("6")
	} else if x == 8 {
		fmt.Println("8")
	} else {
		fmt.Println("else")
	}
}
func ninja_3_8() {
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}
func ninja_3_9() {
	favSport := "alfa"
	switch favSport {
	case "alfa":
		fmt.Println(favSport)
	}
}
func ninja_3_10() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
func ninja_4_1() {
	x := [5]int{42, 43, 44, 45, 46}

	fmt.Println(x)
}
func ninja_4_2() {

}
func ninja_4_3() {

}
func ninja_4_4() {

}
func ninja_4_5() {

}
func ninja_4_6() {

}
func ninja_4_7() {

}
func ninja_4_8() {

}
func ninja_4_9() {

}
func ninja_4_10() {

}
