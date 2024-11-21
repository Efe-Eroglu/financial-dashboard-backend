package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Global DB bağlantısı
var DB *sqlx.DB

// ConnectDB: Veritabanına bağlanır ve global DB değişkenini atar
func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız oldu: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız oldu: %v", err)
	}

	log.Println("Veritabanı bağlantısı başarılı")
}
