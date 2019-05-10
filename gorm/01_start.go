package main

import (
	"gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connected database")
	}
	defer db.Close()

	//migrate the schema
	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{Code: "l1212", Price: 1000})

	//Read
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "l1212")

	//update
	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}
