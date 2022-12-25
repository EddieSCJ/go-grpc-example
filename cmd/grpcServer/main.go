package main

import (
	"database/sql"
	"github.com/EddieSCJ/go-grpc-example/internals/database"
	"github.com/EddieSCJ/go-grpc-example/internals/pb"
	"github.com/EddieSCJ/go-grpc-example/internals/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer closeDB(db)

	categoryDB := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		closeDB(db)
	}
}
