package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(
	host string,
	user string,
	password string,
	dbname string,
) {
	// Implementation for connecting to the database goes here
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, dbname,
	)

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// // Perform auto-migrations
	// err = db.AutoMigrate(
	// 	&models.User{},
	// 	&models.MealPlan{},
	// 	&models.MealLog{},
	// 	&models.ReminderLog{},
	// )
	// if err != nil {
	// 	log.Fatalf("❌ Failed to migrate tables: %v", err)
	// }

	// DB = db
	fmt.Println(" Database connected successfully!")
}
