package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
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

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız oldu : %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız oldu : %v", err)
	}

	log.Println("Veritabanı bağlantısı başarılı")
	return db
}
