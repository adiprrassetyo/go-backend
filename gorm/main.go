package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := "host=localhost user=<user_postgres> password=<password_postgres> dbname=<database_postgres> port=<port_postgres> sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Aditira", Age: 23})

	// Read
	// SELECT * FROM "users" WHERE "users"."id" = 1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1
	var user User
	db.First(&user, 1)                     // temukan user dengan menggunakan primary key dan simpan di variabel user
	db.First(&user, "name = ?", "Aditira") // temukan user dengan nama Aditira

	fmt.Println(user)
}
