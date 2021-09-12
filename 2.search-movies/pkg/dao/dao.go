package dao

import "search-movies/pkg/model"

type DB interface {
	InsertLog(log model.Log) error

	MigrateDB(models ...interface{})
	Health() error
}
