package main

import (
	"os"
	"log"
	"fmt"

	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func main() {
	db := setupDB()

	router := gin.Default()

	server := NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":9000", router))
}

func setupDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	if err := db.AutoMigrate(&User{}, &Disease{}, &Med{}, &Clinic{}); err != nil {
		panic(err)
	}

	return db
}