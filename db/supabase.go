package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("❌ DB_URL environment variable not set")
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database ping failed:", err)
	}

	log.Println("✅ Connected to Supabase database.")
}
