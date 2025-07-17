package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Replace with your actual Supabase connection string
	connStr := "postgres://postgres:<your-password>@db.ylvkvzitlfurfxtfclcb.supabase.co:5432/postgres?sslmode=require"

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
