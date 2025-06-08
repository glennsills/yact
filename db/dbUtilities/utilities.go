package dbUtilities

import (
	"log"

	"github.com/glennsills/yact/config"
	usersData "github.com/glennsills/yact/db/usersData"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDb() {
	db, err := gorm.Open(postgres.Open(config.App.Dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	err = db.AutoMigrate(&usersData.User{})
	if err != nil {
		log.Fatal(err)
	}

}
