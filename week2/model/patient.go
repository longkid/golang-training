package model

import (
	"time"
)

// Define an enum
type Gender int

const (
	Undefined Gender = iota
	Male
	Female
)

type Person struct {
	ID       uint
	Fullname string
	Address  string
	Birthday time.Time
	Gender   Gender
	age      int
}

// Constructor for Person
func NewPerson(fullname string, address string, gender Gender) *Person {
	p := Person{
		Fullname: fullname,
		Address:  address,
		Birthday: time.Date(1985, time.October, 21, 0, 0, 0, 0, time.UTC),
		Gender:   gender,
		age:      -1,
	}
	return &p
}

func (p *Person) GetAge() int {
	if p.age != -1 {
		return p.age
	}

	p.age = time.Now().Year() - p.Birthday.Year()
	return p.age
}

// TableName overrides the table name used by User to `patients`
func (Person) TableName() string {
	return "patients"
}

// Define Patient
type Patient struct {
	Person
	InfectedDate time.Time
}

func NewPatient(fullname string, address string, gender Gender, infectedDate time.Time) *Patient {
	p := Patient{
		Person{
			Fullname: fullname,
			Address:  address,
			Birthday: time.Date(1985, time.October, 21, 0, 0, 0, 0, time.UTC),
			Gender:   gender,
			age:      -1,
		},
		infectedDate,
	}
	return &p
}
