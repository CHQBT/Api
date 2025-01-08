package main

import (
	"context"
	"fmt"
	"log"
	"milliy/config"
	pb "milliy/generated/api"
	"milliy/middleware"
	"milliy/service"
	"milliy/storage/postgres"
	"milliy/upload"
	"net"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	store, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}

	uploader, err := upload.NewMinioUploader()
	if err != nil {
		log.Fatal(err)
	}

	storage := postgres.NewPostgresStorage(store)
	twitSvc := service.NewTwitService(storage, uploader)
	userSvc := service.NewUserService(storage)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterTwitServiceServer(grpcServer, twitSvc)
	pb.RegisterUserServiceServer(grpcServer, userSvc)

	go func() {
		log.Printf("gRPC server listening on :9090")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server with gRPC-Gateway
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterTwitServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts); err != nil {
		log.Fatalf("Failed to register twit gateway: %v", err)
	}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts); err != nil {
		log.Fatalf("Failed to register user gateway: %v", err)
	}

	// Serve Swagger UI
	fs := http.FileServer(http.Dir("doc/swagger"))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	http.Handle("/", mux)

	// Initialize Casbin enforcer
	enforcer, err := casbin.NewEnforcer("casbin/model.conf", "casbin/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Use middleware
	handler := cors.Default().Handler(
		middleware.AuthMiddleware(enforcer)(
			http.DefaultServeMux,
		),
	)

	// Start HTTP server
	httpPort := fmt.Sprintf(":%s", cfg.Server.HTTP_PORT)
	log.Printf("HTTP server listening on %s", httpPort)
	log.Printf("Swagger UI available at http://localhost%s/swagger/", httpPort)
	if err := http.ListenAndServe(httpPort, handler); err != nil {
		log.Fatal(err)
	}
}
