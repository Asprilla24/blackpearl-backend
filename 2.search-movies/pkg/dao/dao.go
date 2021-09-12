package dao

import "search-movies/pkg/model"

// RequiredDBVersion is database version required by the service to operate
// If you need to make changes to database model here, please create db migration
// in {project_dir}/migration (details : https://github.com/golang-migrate/migrate).
// Run migration script in {project_dir}/migration.sh afterward
const RequiredDBVersion = 1

// SchemaMigration is database model of schema migration
type SchemaMigration struct {
	Version int
	Dirty   bool
}

type DB interface {
	InsertLog(log model.Log) error

	MigrateDB(models ...interface{})
	Health() error
}
