package postgres_test

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"search-movies/pkg/config"
	"search-movies/pkg/dao"
	"search-movies/pkg/dao/postgres"
	"search-movies/pkg/model"

	"github.com/caarlos0/env"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	confTest *config.Config
	dbConn   *gorm.DB
	db       dao.DB

	accessLog = model.Log{
		URL:            "http://testing-log/search",
		ResponseStatus: http.StatusOK,
		CreatedAt:      time.Now(),
	}
)

// nolint : dupl
func TestMain(m *testing.M) {
	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "search-movies", log.DefaultTimestampUTC)

	confTest = &config.Config{}
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)
		for _, val := range confTest.HelpDocs() {
			fmt.Println(val)
		}
		fmt.Println("")
		flag.PrintDefaults()
	}
	flag.Parse()

	err := env.Parse(confTest)
	if err != nil {
		logger.Log("ENV", "Parse", "err", err)
		return
	}

	dbConn, err = dao.NewPostgres("postgres", confTest)
	if err != nil {
		logger.Log("DAO", "NewPostgres", "err", err)
		return
	}
	defer dbConn.Close() // nolint : errcheck, used in defer

	db = postgres.NewDB(dbConn)
	db.MigrateDB(&model.Log{})

	os.Exit(m.Run())
}

func TestDB_InsertLog(t *testing.T) {
	err := db.InsertLog(accessLog)

	var actualLog model.Log
	dbConn.First(&actualLog)

	assert.Nil(t, err, "InsertLog should return nil error")
	assert.Equal(t, accessLog.URL, actualLog.URL,
		"saved and retrieved log's URL should be equal")
	assert.Equal(t, accessLog.ResponseStatus, actualLog.ResponseStatus,
		"saved and retrieved log's ResponseStatus should be equal")

	dbConn.Delete(&model.Log{})
}
