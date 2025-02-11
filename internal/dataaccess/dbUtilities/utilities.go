package dbUtilities

import (
	"log"

	"github.com/glennsills/yact/internal/dataaccess/users_db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToSQLite(dbfile string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbfile), config)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&users_db.User{})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
