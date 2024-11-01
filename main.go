package main

import (
	"context"
	"log"
	"net"
	"time"

	accessorySvc "lostark/internal/accessory"
	tasks "lostark/internal/accessory/task"
	accessoryPb "lostark/proto/accessory"

	"google.golang.org/grpc"
)

func main() {

	// dbConfig := database.Config{
	// 	Host:     "localhost",
	// 	Port:     3306,
	// 	User:     "user",
	// 	Password: "password",
	// 	Database: "accessories_db",
	// }

	// mysql, err := database.NewMySQL(dbConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer mysql.Close()

	//scheduler
	taskManager := tasks.NewTaskManager()
	accessoryTask := tasks.NewAccessoryPriceTask(taskManager)
	go accessoryTask.RunImmediately()
	if err := accessoryTask.Schedule(); err != nil {
		log.Fatalf("Failed to schedule accessory price task: %v", err)
	}
	taskManager.Start()

	//grpc
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	accessoryService := accessorySvc.NewService()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	accessoryPb.RegisterAccessoryServiceServer(server, accessoryService)

	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(start)

	log.Printf("%s / %s / %s",
		info.FullMethod,
		start.Format("2006-01-02 15:04:05"),
		duration)

	return resp, err
}
