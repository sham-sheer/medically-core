package main

import (
	"log"
	"database/sql"
	"fmt"
	"os"
	"context"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cenkalti/backoff/v4"
	"github.com/cockroachdb/cockroach-go/v2/crdb"
)


func main() {
	db, err := initStore()
	if err != nil {
		log.Fatalf("Failed to initialise the store: %s", err)
	}
	defer db.Close()

	r := gin.Default()
	log.Println("Starting Server...")
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "medically-core at your service!!",
		})
	})

	r.GET("/count", func(c *gin.Context) {
		countHandler(db, c)
	})

	r.POST("/add-message", func(c *gin.Context) {
		addMessageHandler(db, c)
	})
	r.Run()
}

func countHandler(db *sql.DB, c *gin.Context) {
	r, err := countRecords(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{
		"count": r,
	})
}

type Message struct {
	Value string `json:"value"`
}
func addMessageHandler(db *sql.DB, c *gin.Context) {

	m := &Message{}
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Println("Adding to db: ", m.Value)
	err := crdb.ExecuteTx(context.Background(), db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"INSERT INTO message (value) VALUES ($1) ON CONFLICT (value) DO UPDATE SET value = excluded.value",
				m.Value,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return nil
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": m.Value})
}

func initStore() (*sql.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	var (
		db  *sql.DB
		err error
	)
	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS message (value STRING PRIMARY KEY)"); err != nil {
		return nil, err
	}

	return db, nil
}

func countRecords(db *sql.DB) (int, error) {

	rows, err := db.Query("SELECT COUNT(*) FROM message")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
		rows.Close()
	}

	return count, nil
}