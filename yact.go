package main

import (
	"log"
	"net/http"

	"github.com/glennsills/yact/ui"
)

func main() {
	ui.AssignHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {

// 	dsn := "host=database-1.c7im4k02oirc.us-east-2.rds.amazonaws.com dbname=yact port=5432 user=postgres password=SuperSecret dbname=yact sslmode=allow"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sqlDB.Close()

// 	err = db.AutoMigrate(&users_db.User{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create a user
// 	dob := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	newUser := &users_db.User{
// 		UserName:  "jdoe",
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Email:     "john.doe@example.com",
// 		Dob:       &dob,
// 	}

// 	repo := users_db.NewUserRepository(db)

// 	err = repo.CreateUser(newUser)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = repo.DeleteUserById(newUser.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("%v", newUser)

// 	log.Default().Println("here")
// }
