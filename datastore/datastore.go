package datastore

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" //  postgreSQL initialization
)

// Db is a global variable.
var Db *sql.DB

func init() {
	// Load env config
	err := godotenv.Load("chatter/.env")
	if err != nil {
		log.Fatalf("Failed loading .env config: %v", err)
	}
	dbname, sslmode := os.Getenv("dbname"), os.Getenv("sslmode")
	Db, err = sql.Open("postgres", dbname+" "+sslmode)
	if err != nil {
		log.Fatalf("Failed opening sql driver: %v", err)
	}
}

// Encrypt hashes plain text with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
