package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maspratama/go-patungan/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConncetToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connect := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("Connection to db failed!")
	}

	DB = db
	fmt.Println("Connection to db success!")

	AutoMigrate(db)
}
func AutoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate(
		&entity.Users{},
		//&entity.Transactions{},
		//&entity.Campaigns{},
		//&entity.CampaignImages{},
	)
	if err != nil {
		return
	}
}
