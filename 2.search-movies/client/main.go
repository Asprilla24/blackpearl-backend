package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"search-movies/api/v1/searchMovies"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:9080", "The server address in the format of host:port")
)

func runSearchMovies(client searchMovies.SearchMoviesClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Search(ctx, &searchMovies.SearchRequest{Pagination: 2, SearchWord: "batman"})
	if err != nil {
		fmt.Printf("%v.Search(ctx) = %v, %v: ", client, stream, err)
	}
	fmt.Println(stream)
}

func runHealthCheck(client searchMovies.SearchMoviesClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.HealthCheck(ctx, &searchMovies.HealthCheckRequest{})
	if err != nil {
		fmt.Printf("%v.HealthCheck(ctx) = %v, %v: ", client, stream, err)
	}
	fmt.Println(stream)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Printf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := searchMovies.NewSearchMoviesClient(conn)

	runSearchMovies(client)
	runHealthCheck(client)
}
