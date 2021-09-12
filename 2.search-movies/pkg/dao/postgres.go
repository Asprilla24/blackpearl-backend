package dao

import (
	"fmt"
	"time"

	"search-movies/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	"github.com/pkg/errors"
)

const (
	dbTimeoutConnection = 30 * time.Second
)

// Config define postgres config
type PostgresConfig struct {
	Host       string
	Username   string
	Password   string
	Name       string
	Port       int
	SSLEnabled bool
	LogMode    bool
}

// New creates new instance of DAO for database operations.
// dbDriver is database driver name,
// see http://doc.gorm.io/database.html#connecting-to-a-database for available database driver
func NewPostgres(dbDriver string, conf *config.Config) (*gorm.DB, error) {
	sslMode := "" // omit sslmode so it set to default (enabled)
	if !conf.DBSSLEnabled {
		sslMode = "sslmode=disable"
	}
	connectionConfig := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=5 TimeZone=Asia/Jakarta "+sslMode,
		conf.DBHost, conf.DBPort, conf.DBName, conf.DBUsername, conf.DBPassword)

	timeout := time.Now().Add(dbTimeoutConnection)
	var postgresORM *gorm.DB
	var err error
	retryCounter := 0
	for time.Now().Before(timeout) {
		postgresORM, err = gorm.Open(dbDriver, connectionConfig)
		if err == nil {
			break
		}
		retryCounter++
	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to connect to database: timeout: %v", err)
	}
	if postgresORM == nil {
		return nil, errors.Wrapf(err, "unable to initiate DAO: %v", err)
	}
	return postgresORM, nil
}
