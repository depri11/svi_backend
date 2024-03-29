package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sharing_vasion_indonesia/article_service/database"
	"sharing_vasion_indonesia/article_service/internal/delivery"
	"sharing_vasion_indonesia/article_service/internal/repository"
	"sharing_vasion_indonesia/article_service/internal/usecase"
	article_proto "sharing_vasion_indonesia/pkg/proto"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := ":9001"

	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	repo := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repo)
	delivery := delivery.NewDelivery(usecase)

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
		Timeout:           15 * time.Second,
		MaxConnectionAge:  5 * time.Minute,
		Time:              5 * time.Minute,
	}))

	article_proto.RegisterArticleServiceServer(server, delivery)

	go func() {
		err := server.Serve(l)
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Printf("GRPC Server Article is listening on port: %v", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server GRPC Article Exited Properly")
}
