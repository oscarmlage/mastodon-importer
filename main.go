package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"

	"mastodon-importer/models"
)

func main() {
	fmt.Println("Mastodon importer")

	// Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln(err)
	}

	// Load JSON file
	content, err := ioutil.ReadFile("./content/sample.json")
	if err != nil {
		log.Panicln(err)
	}

	var payload models.JSONData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Panicln(err)
	}

	for _, curr := range payload.OrderedItems {
		log.Printf("item: %s\n", curr)
	}

	// Connect DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("host"), os.Getenv("user"), os.Getenv("password"),
		os.Getenv("dbname"), os.Getenv("port"), os.Getenv("sslmode"),
		os.Getenv("timezone"))
	log.Printf("dsn: %s", dsn)

	os.Exit(0)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	// Create new Status (toot)
	var toot models.Status
	toot.Text = "Hello world"
	toot.Created_at = time.Now()
	toot.Updated_at = time.Now()
	toot.Sensitive = "f"
	toot.Visibility = 0
	toot.Local = "t"
	toot.Account_id = 108257603645889535
	toot.Application_id = 1

	// Save it
	tx := db.Create(&toot)
	if tx.Error != nil {
		log.Panicln(tx.Error)
	}
}
