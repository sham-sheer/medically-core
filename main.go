package main

import (
	"flag"
	"log"
	"fmt"

	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	addr = flag.String("addr", "addr", "addr to connect to database")
)

func main() {
	flag.Parse()

	db := setupDB(*addr)

	router := gin.Default()

	server := NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":6543", router))
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(addr))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	if err := db.AutoMigrate(&User{}, &Disease{}, &Med{}); err != nil {
		panic(err)
	}

	return db
}