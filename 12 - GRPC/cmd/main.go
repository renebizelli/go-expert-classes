package main

import (
	"database/sql"
	"net"
	"renebizelli/grpc/internal/database"
	"renebizelli/grpc/internal/pb"
	"renebizelli/grpc/internal/services"

	_ "github.com/mattn/go-sqlite3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := services.NewCategoryService(categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService) // RegisterCategoryServiceServer gerado no arquivo course_category_grpc.pb.go
	reflection.Register(grpcServer)                               // Para o Evans, client.

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
