package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	ninja_9_6()
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
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x)
}
func ninja_4_2() {
	x := []int{42, 43, 44, 45, 46, 5, 1, 3, 79, 346}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x)
}
func ninja_4_3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[3:5])
}
func ninja_4_4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
func ninja_4_5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	z := []int{}
	z = append(x[:3], x[6:]...)
	fmt.Println(z)
}
func ninja_4_6() {
	x := make([]string, 50, 50)
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	for i, v := range states {
		x[i] = v
	}
	fmt.Println((cap(x)))
	fmt.Println(len(x))
	fmt.Println(x)
}
func ninja_4_7() {

	x1 := []string{"raz", "dwa", "czy"}
	x2 := []string{"aaa", "sss", "ddd"}
	x := [][]string{x1, x2}
	fmt.Println(x)
	for i1, v1 := range x {
		for i2, v2 := range v1 {
			fmt.Println("wiersz:", i1, " wpis:", i2, " wartość:", v2)
		}
	}
}
func ninja_4_8() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
func ninja_4_9() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["Romek"] = []string{"kobity", "dinozaury"}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
func ninja_4_10() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["Romek"] = []string{"kobity", "dinozaury"}
	if _, ok := m["Romek"]; ok {
		fmt.Println(`usuwamy go! `)
		delete(m, "Romek")
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
func ninja_5_1() {
	type person struct {
		first string
		last  string
	}
	p1 := person{
		first: "Alfa",
		last:  "Bravo",
	}
	p2 := person{
		first: "Charlie",
		last:  "Delta",
	}
	for i, v := range p1.first {
		fmt.Println(i, "_", v)
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
func ninja_5_2() {
	type person struct {
		first string
		last  string
	}
	p1 := person{
		first: "Alfa",
		last:  "Bravo",
	}
	p2 := person{
		first: "Charlie",
		last:  "Delta",
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
func ninja_6_1() {
	fmt.Println(bar61("alfa", 42))
	fmt.Println(foo61(7))
}
func foo61(x int) int {
	return x + 2
}
func bar61(s string, x int) (string, int) {

	return (s + "X"), x - 1
}
func ninja_6_2() {
	fmt.Println(foo62(7, 5, 3))
	x := []int{7, 5, 3, 2}
	x = append(x, 7)
	fmt.Println(bar62(x))
}
func foo62(x ...int) int {
	y := 0
	for _, v := range x {
		y += v
	}
	return y
}
func bar62(x []int) int {
	y := 0
	for _, v := range x {
		y += v
	}
	return y
}
func ninja_6_3() {
	fmt.Println("alfa")
	defer fmt.Println("beta")
	fmt.Println("gama")
}

type person64 struct {
	first string
	last  string
	age   int
}

func (p person64) speak64() {
	fmt.Println("I'm", p.first, p.last, "and I'm", p.age, "y. old")
}
func ninja_6_4() {
	p1 := person64{
		first: "Jack",
		last:  "Sparrow",
		age:   99,
	}
	p1.speak64()
}

type square65 struct {
	a float64
}
type circle65 struct {
	a float64
}

func (s square65) area65() float64 {
	return float64(s.a * s.a)
}
func (c circle65) area65() float64 {
	return float64(c.a * c.a * math.Pi)
}
func ninja_6_5() {
	s1 := square65{
		a: (5.4363456),
	}
	c1 := circle65{
		a: (4.762),
	}
	fmt.Println(s1.area65())
	fmt.Println(c1.area65())
}
func ninja_6_6() {
	x := func() int {
		return 7
	}
	y := x()
	fmt.Printf("%t____%v\n", x, x)
	fmt.Println(y)
}
func retval68() func() int {
	fmt.Println("retval68")
	x := 0
	return func() int {
		x += 2
		return x
	}
}
func ninja_6_8() {
	x := retval68()
	fmt.Println(x())
	fmt.Println(x())
}
func ninja_6_9() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo69(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}
func foo69(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

type user81 struct {
	First string
	Age   int
}

func ninja_8_1() {
	u1 := user81{
		First: "James",
		Age:   32,
	}

	u2 := user81{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user81{
		First: "M",
		Age:   54,
	}

	users := []user81{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	var users2 []user81
	err2 := json.Unmarshal(bs, &users2)
	if err2 != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(users2)

}

type user83 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func ninja_8_3() {
	u1 := user83{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user83{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user83{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user83{u1, u2, u3}

	fmt.Println(users, "\n")

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}
}

var wg91 sync.WaitGroup

func ninja_9_1() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg91.Add(2)
	go func() {
		fmt.Println("jeden")
		wg91.Done()
	}()
	go func() {
		fmt.Println("dwa")
		wg91.Done()
	}()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg91.Wait()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

type person92 struct {
	First string
}
type human92 interface {
	speak92()
}

func (p *person92) speak92() {
	fmt.Println("hello")
}
func saySomething92(h human92) {
	h.speak92()
}
func ninja_9_2() {
	p1 := person92{
		First: "Roman",
	}
	fmt.Println(p1)
	saySomething92(&p1)
}
func ninja_9_3() {
	fmt.Println("CPUs:\n", runtime.NumCPU())
	fmt.Println("GoRoutiness:", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Printf("Counter:%v", counter)
		fmt.Printf("\tRoutines: %v\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutiness:", runtime.NumGoroutine())
	fmt.Println("koniec: ", counter)
}
func ninja_9_4() {
	fmt.Println("CPUs:\n", runtime.NumCPU())
	fmt.Println("GoRoutiness:", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Printf("Counter:%v", counter)
		fmt.Printf("\tRoutines: %v\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutiness:", runtime.NumGoroutine())
	fmt.Println("koniec: ", counter)
}
func ninja_9_5() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
func ninja_9_6() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
