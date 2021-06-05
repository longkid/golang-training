package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longkid/golang-training/week2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Connect DB
	dsn := "root:secret@tcp(127.0.0.1:3306)/covid?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Implement REST APIs
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/patients/:id", func(c *gin.Context) {
		id := c.Param("id")
		var patient model.Patient
		db.First(&patient, id)
		// fmt.Printf("Get patient with id '%v': %v\n", id, patient)
		// c.String(200, "GET patient: SUCCESS")
		c.JSON(200, patient)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
