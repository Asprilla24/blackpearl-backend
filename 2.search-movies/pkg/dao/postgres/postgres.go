package postgres

import (
	"search-movies/pkg/dao"
	"search-movies/pkg/model"

	"github.com/jinzhu/gorm"
)

type DB struct {
	conn *gorm.DB
}

func NewDB(conn *gorm.DB) dao.DB {
	return &DB{conn: conn}
}

func (db *DB) InsertLog(log model.Log) error {
	return db.conn.Create(&log).Error
}

//MigrateDB for migrate an models
func (db *DB) MigrateDB(models ...interface{}) {
	for _, model := range models {
		db.conn.AutoMigrate(model)
	}
}

// Health check ping postgres
func (db *DB) Health() error {
	return db.conn.DB().Ping()
}
