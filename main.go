package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	accessorySvc "lostark/internal/accessory"
	"lostark/internal/accessory/task"
	accessoryPb "lostark/proto/accessory"

	"github.com/hibiken/asynq"
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

	// accessoryRepo := models.NewAccessoryRepository(mysql.DB())

	redisAddr := "localhost:6379"

	scheduler, err := task.NewScheduler(redisAddr)
	if err != nil {
		log.Fatalf("Failed to create scheduler: %v", err)
	}

	if err := scheduler.RegisterPriceUpdateTask(); err != nil {
		log.Fatalf("Failed to register price update task: %v", err)
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 1,
			Queues: map[string]int{
				"default": 1,
			},
		},
	)

	// handler := task.NewTaskHandler(accessoryRepo)
	mux := asynq.NewServeMux()
	// handler.RegisterHandlers(mux)

	go func() {
		if err := scheduler.Start(); err != nil {
			log.Fatalf("Failed to start scheduler: %v", err)
		}
	}()

	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// 종료 시그널 처리
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// 정상 종료
	scheduler.Shutdown()
	srv.Shutdown()

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
