package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint // uint : 価格は常に非負であるため、この型が適しています
}

func main() {
	dsn := "host=postgres_db user=root password=password dbname=dev port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Product{})
	db.Create(&Product{
		Code:  "D2Micro",
		Price: 10,
	})

	var product Product
	db.First(&product, 1)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": product,
		})
	})

	router.Run(":8080")
}