package main

import (
	"log"

	"github.com/JFMajer/microservices/order/config"
	"github.com/JFMajer/microservices/order/internal/adapters/db"
	"github.com/JFMajer/microservices/order/internal/adapters/grpc"
	"github.com/JFMajer/microservices/order/internal/application/core/api"
)

func main() {
	// Log start of program
	log.Println("Starting program...")

	// Attempt to create the DB adapter
	log.Println("Connecting to database...")
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	log.Println("Successfully connected to the database.")

	// Create the application with the dbAdapter
	log.Println("Creating application instance...")
	application := api.NewApplication(dbAdapter)

	// Create the grpc adapter
	log.Println("Creating gRPC adapter...")
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	// Log when gRPC server starts
	log.Println("Starting gRPC server...")
	grpcAdapter.Run()

	// Log when the program finishes (in case it exits gracefully)
	log.Println("Program has finished running.")
}
