package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model

	ID    string
	Title string
}

func main() {
	const dsn = "host=localhost user=postgres password=root dbname=fghub-db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Game{})

	// Create
	db.Create(&Game{ID: "CPNSF5", Title: "Street Fighter V"})

	// Read
	var product Game
	// db.First(&product, 1)                  // find product with integer primary key
	db.First(&product, "ID = ?", "CPNSF5") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Title", "Test")
	// Update - update multiple fields
	db.Model(&product).Updates(Game{ID: "CPNSF5", Title: "Street Fighter 5"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"ID": "CPNSF5", "Title": "SFive"})

	// Delete - delete product
	// db.Delete(&product, 1)
}
