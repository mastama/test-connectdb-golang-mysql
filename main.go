package main

import (
	"patungan-be/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
	"log"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:tams4596@tcp(127.0.0.1:3306)/patungan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection to database is success!")

	var users []entity.User
	db.Find(&users)

	// untuk menampilkan data
	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("====================")
	}
}
