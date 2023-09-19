package main

import (
	"fmt"
)

var I = "12" // Shadow

/* func main() {

	//Variables
	var i int = 1
	// var i string
	// i = "20"
	// i = "90"
	// var j int = 20
	// var j float32 = float32(i)
	var j float32 = float32(i)
	s := 12.12
	// var k string
	// var foo string = string(i)
	var foo int = int(i)
	var foo1 string = strconv.Itoa(i)
	var bar bool = true
	fmt.Printf("%v, %T\n", I, I)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", s, s)
	// fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", foo, foo)
	fmt.Printf("%v, %T\n", foo1, foo1)
	fmt.Printf("%v, %T\n", bar, bar)
	fmt.Println("==========================START PRIMITIVES===================")

	//Primitives
	// var b complex128 = 1 + 2i
	var b complex128 = complex(1, 2) // Dynamic Create Complex number
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", real(b), real(b))
	fmt.Printf("%v, %T\n", imag(b), imag(b))
	fmt.Println("==========================START PRIMITIVES Arthmertic operators  with complex number===================")

	var c complex128 = complex(2, 3)
	d := 1 + 6i
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	// fmt.Println(i)

	fmt.Println("==========================START PRIMITIVES with string===================")
	e := "This is string"
	e1 := e[0]
	e2 := string(e[0])
	e3 := []byte(e)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", e1, e1) // retun byte
	fmt.Printf("%v, %T\n", e2, e2) // retun updated values
	fmt.Printf("%v, %T\n", e3, e3) // retun string Bytes

	f := 'T'                     // Rune Type
	fmt.Printf("%v, %T\n", f, f) // retun string Bytes
} */

/* const (
	i = iota + 1
	_
	j
	_
	k
)

func main() {
	fmt.Println("===================Constaints========================")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
} */

func main() {
	fmt.Println("===================Arrays========================")
	var fixArr [3]int = [3]int{1, 2, 3}
	a2 := &fixArr
	/* fmt.Printf("%v\n", fixArr)
	dynamicArr := [...]int{1, 2, 4, 2}
	// dynamicArr[4] = 12
	fmt.Printf("dynamicArr:  %v\n", dynamicArr[3])
	fmt.Printf("dynamicArr length:  %v", len(dynamicArr)) */
	fmt.Printf("fixarr: %v\n", fixArr)
	fmt.Printf("assign: %v\n", a2)
	fixArr[0] = 100

	fmt.Printf("fixarr1: %v\n", fixArr)
	fmt.Printf("assign1: %v\n", a2)

	newArr := [...]int{1, 27, 45, 34, 5343, 232, 43, 5, 45, 4, 5, 45, 4, 45, 5, 4, 54}
	a := newArr[:]
	b := newArr[2:]
	c := newArr[:5]
	d := newArr[3:6]
	fmt.Printf("A: %v\n", a)
	fmt.Printf("B: %v\n", b)
	fmt.Printf("C: %v\n", c)
	fmt.Printf("D: %v\n", d)

}
