package main

import (
	"fmt"
	"strconv"
)

var I = "12" // Shadow

func main() {
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
	// fmt.Println(i)
}
