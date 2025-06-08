package lib

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm Object Database
var DB *gorm.DB

// Connect
func ConnectToAuthDatabase() {
	var err error

	host := os.Getenv("AUTH_DB_HOST")
	name := os.Getenv("AUTH_DB_NAME")
	pass := os.Getenv("AUTH_DB_PASS")
	port := os.Getenv("AUTH_DB_PORT")
	user := os.Getenv("AUTH_DB_USER")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port,
	)

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		log.Infof("connecting to DB (%s:%s/%s), attempt %d/%d", host, port, name, attempts, maxAttempts)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Infof("successfully connected to %s:%s/%s", host, port, name)
			return
		}

		log.Warnf("database connection failed: %v", err)

		time.Sleep(1 * time.Second) // wait before retry
	}

	log.Fatalf("could not connect to database after %d attempts: %v", maxAttempts, err)
}
