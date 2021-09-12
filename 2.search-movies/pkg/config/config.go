package config

import (
	"fmt"
	"reflect"
)

// Config specifies configurable options through env vars
// nolint:lll
type Config struct {
	HTTPPort string `env:"SERVER_HTTP_PORT" envDocs:"The http port which the service will listen to" envDefault:"8081"`
	GRPCPort string `env:"SERVER_GRPC_PORT" envDocs:"The grpc port which the service will listen to" envDefault:"9080"`

	OMDBKey string `env:"OMDB_KEY,required" envDocs:"The OMDB API Key"`

	// database configurations
	DBHost       string `env:"DB_HOST,required" envDocs:"Database host server"`
	DBPort       int    `env:"DB_PORT,required" envDocs:"Database port server" envDefault:"5432"`
	DBName       string `env:"DB_NAME,required" envDocs:"Database name"`
	DBUsername   string `env:"DB_USERNAME,required" envDocs:"Database username"`
	DBPassword   string `env:"DB_PASSWORD,required" envDocs:"Database password"`
	DBSSLEnabled bool   `env:"DB_SSL_ENABLED" envDocs:"Use SSL for database connection" envDefault:"true"`
}

// HelpDocs returns documentation of Config based on field tags
func (envVar Config) HelpDocs() []string {
	reflectEnvVar := reflect.TypeOf(envVar)
	doc := make([]string, 1+reflectEnvVar.NumField())
	doc[0] = "Environment variables config:"
	for i := 0; i < reflectEnvVar.NumField(); i++ {
		field := reflectEnvVar.Field(i)
		envName := field.Tag.Get("env")
		envDefault := field.Tag.Get("envDefault")
		envDocs := field.Tag.Get("envDocs")
		doc[i+1] = fmt.Sprintf("  %v\t %v (default: %v)", envName, envDocs, envDefault)
	}
	return doc
}
