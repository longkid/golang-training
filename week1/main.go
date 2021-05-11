package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/longkid/golang-training/week1/helper"
	"github.com/longkid/golang-training/week1/model"
)

func main() {
	pointers()
}

func main2() {
	fmt.Println("Hello world")
	fmt.Println("Sum of 1, 2, 3, 4:", helper.Sum(1, 2, 3, 4))
	p := model.Person{
		Name:          "Phu",
		Year_of_birth: 1986,
	}
	data, _ := json.Marshal(p)
	fmt.Println("JSON marshal:", string(data))

	s := "100a000"
	fmt.Println("len of s:", len(s))

	var n int
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of %s is %d\n", s, n)
}

func pointers() {
	var n int = 10
	fmt.Printf("Value of n = %v\n", n)
	fmt.Printf("Address of n = %v\n", &n)
	var p *int = &n
	fmt.Printf("Value of p = %v\n", p)
	fmt.Printf("Address of p = %v\n", &p)
	fmt.Printf("Value of *p = %v\n", *p)

	// change n -> *p is changed too
	n = -9
	fmt.Printf("Value of *p = %v\n", *p)

	// change *p -> n is changed too
	*p = 100
	fmt.Printf("Value of n = %v\n", n)

	// p points to m -> *p will contain value of m
	m := 99
	p = &m
	fmt.Printf("Value of *p = %v\n", *p)

	// change *p -> n is not changed
	*p = 1111111111
	fmt.Printf("Value of n = %v\n", n)
}
