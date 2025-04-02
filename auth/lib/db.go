package lib

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm Object Database
var DB *gorm.DB

// Connect
func ConnectToAuthDatabase() {
	var err error

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	pass := os.Getenv("DB_PASS")

	log.Infof("Connecting to %s:%s/%s", host, name, port)

	dsn := fmt.Sprintf(
		"host=%s user=postgres password=%s dbname=%s port=%s sslmode=disable application_name=auth",
		host,
		pass,
		name,
		port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to connect to %s:%s/%s: %s", host, port, name, err))
	} else {
		log.Infof("%s:%s/%s connection successful", host, port, name)
	}

}
