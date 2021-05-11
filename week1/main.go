package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/longkid/golang-training/week1/helper"
	"github.com/longkid/golang-training/week1/model"
	"gopkg.in/yaml.v3"
)

func main() {
	// main2()
	// pointers()
	convertCsvToYml()
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

type Person struct {
	Id          int
	Name        string
	YearOfBirth int
}

func convertCsvToYml() {
	inFile, err := os.Open("week1/data/input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inFile.Close()
	fmt.Println("Input:", inFile.Name())

	scanner := bufio.NewScanner(inFile)

	var persons []Person
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "ID" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		yearOfBirth, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		persons = append(persons, Person{
			Id:          id,
			Name:        parts[1],
			YearOfBirth: yearOfBirth,
		})
	}

	out, _ := yaml.Marshal(persons)
	ymlContent := string(out)
	// fmt.Println(ymlContent)

	// Write out to file output.yml
	outFile, err := os.Create("week1/data/output.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	_, err = outFile.WriteString(ymlContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Output:", outFile.Name())
}
