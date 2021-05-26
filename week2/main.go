package main

import (
	"fmt"
	"time"

	"github.com/longkid/golang-training/week2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	p := model.NewPerson("Lam Ho", "HCMC", model.Male)
	fmt.Println("person:", *p)
	fmt.Println("1:", p.GetAge())
	fmt.Println("2:", p.GetAge()) // This time, method GetAge() returns p.age imediately

	dsn := "root:secret@tcp(127.0.0.1:3306)/covid?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		fmt.Printf("Could not connect DB: %v\n", err)
		return
	}

	// Ref: https://gorm.io/docs/migration.html
	db.AutoMigrate(&model.Person{})
	db.AutoMigrate(&model.Patient{})

	patient := model.NewPatient("Lam Ho", "HCMC", model.Male, time.Now())
	fmt.Println("patient:", *patient)
	fmt.Printf("Before create in DB: patient.ID = %v\n", patient.ID)
	fmt.Println(patient.Birthday)
	fmt.Println(patient.InfectedDate)
	fmt.Println(patient.GetAge())

	// err = db.Create(&p).Error
	// if err != nil {
	// 	fmt.Printf("Fail to create person: %v\n", err)
	// }
	// fmt.Println("p.ID", p.ID)
	err = db.Create(&patient).Error
	if err != nil {
		fmt.Printf("Fail to create person: %v\n", err)
	}
	fmt.Printf("After create in DB: patient.ID = %v\n", patient.ID)
}
