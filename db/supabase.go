package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found. Continuing with system environment variables.")
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("❌ DB_URL environment variable not set")
	}

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
