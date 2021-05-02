package main

import (
	"log"
	"net"
	"os"

	"golang-training-shop-grpc/pkg/db"
	"golang-training-shop-grpc/product_server/pkg/api"
	pb "golang-training-shop-grpc/proto/go_proto"

	"google.golang.org/grpc"
)

var (
	serverPort = os.Getenv("SERVER_PORT")
	host       = os.Getenv("DB_USERS_HOST")
	port       = os.Getenv("DB_USERS_PORT")
	user       = os.Getenv("DB_USERS_USER")
	dbname     = os.Getenv("DB_USERS_DBNAME")
	password   = os.Getenv("DB_USERS_PASSWORD")
	sslmode    = os.Getenv("DB_USERS_SSL")
)

func init() {
	if serverPort == "" {
		serverPort = ":8080"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "shop"
	}
	if password == "" {
		password = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {

	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, api.NewProductServer(conn))
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
