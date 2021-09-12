package transports_test

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"

	"search-movies/api/v1/searchMovies"
	"search-movies/pkg/config"
	"search-movies/pkg/dao"
	"search-movies/pkg/dao/postgres"
	"search-movies/pkg/endpoints"
	"search-movies/pkg/model"
	"search-movies/pkg/service"
	"search-movies/pkg/transports"

	"github.com/caarlos0/env"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

var (
	dbConnTest     *gorm.DB
	confTest       *config.Config
	dbTest         dao.DB
	serviceTest    service.Service
	httpServerTest http.Handler
	gRpcServerTest searchMovies.SearchMoviesServer
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

	dbConnTest, err = dao.NewPostgres("postgres", confTest)
	if err != nil {
		logger.Log("DAO", "NewPostgres", "err", err)
		return
	}
	defer dbConnTest.Close() // nolint : errcheck, used in defer

	dbTest = postgres.NewDB(dbConnTest)
	dbTest.MigrateDB(&model.Log{})

	serviceTest = service.NewService(dbTest, *confTest)
	eps := endpoints.NewEndpointSet(serviceTest)
	httpServerTest = transports.NewHTTPHandler(eps)
	gRpcServerTest = transports.NewGRPCServer(eps)

	os.Exit(m.Run())
}
