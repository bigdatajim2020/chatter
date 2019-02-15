package datastore

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/williamzion/chatter/logger"

	_ "github.com/lib/pq" //  postgreSQL initialization
)

var (
	// Db is a global variable.
	Db  *sql.DB
	ctx = context.Background()
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // default path is current working directory
	err := viper.ReadInConfig()
	logger.Info.Println(os.Getwd())
	if err != nil {
		log.Fatalf("Failed loading config file: %v", err)
	}

	user, password, dbname, sslmode := viper.GetString("user"), viper.GetString("password"), viper.GetString("dbname"), viper.GetString("sslmode")
	Db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode))
	if err != nil {
		log.Fatalf("Failed opening sql driver: %v", err)
	}

	// Verify database connection.
	if err = Db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
}

// createUUID creates a random UUID with from RFC 4122.
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string, err error) {
	u := new([16]byte)
	_, err = rand.Read(u[:])
	if err != nil {
		return
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt hashes plain text with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
