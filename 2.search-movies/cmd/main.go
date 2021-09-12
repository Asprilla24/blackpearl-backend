package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "search-movies/api/v1/searchMovies"
	"search-movies/pkg/config"
	"search-movies/pkg/dao"
	"search-movies/pkg/dao/postgres"
	"search-movies/pkg/endpoints"
	"search-movies/pkg/model"
	"search-movies/pkg/service"
	"search-movies/pkg/transports"

	"github.com/caarlos0/env"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "search-movies", log.DefaultTimestampUTC)

	conf := config.Config{}
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)
		for _, val := range conf.HelpDocs() {
			fmt.Println(val)
		}
		fmt.Println("")
		flag.PrintDefaults()
	}
	flag.Parse()

	err := env.Parse(&conf)
	if err != nil {
		logger.Log("ENV", "Parse", "err", err)
		return
	}

	dbConn, err := dao.NewPostgres("postgres", &conf)
	if err != nil {
		logger.Log("DAO", "NewPostgres", "err", err)
		return
	}
	defer dbConn.Close() // nolint : errcheck, used in defer

	db := postgres.NewDB(dbConn)
	db.MigrateDB(&model.Log{})

	var (
		service     = service.NewService(db, conf)
		eps         = endpoints.NewEndpointSet(service)
		httpHandler = transports.NewHTTPHandler(eps)
		grpcServer  = transports.NewGRPCServer(eps)
	)

	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpAddr := fmt.Sprintf("localhost:%s", conf.HTTPPort)
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// The gRPC listener mounts the Go kit gRPC server
		grpcAddr := fmt.Sprintf("localhost:%s", conf.GRPCPort)
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("gRPC", "addr", grpcAddr)
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterSearchMoviesServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Log("exit", g.Run())
}
