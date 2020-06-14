package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go"
	"github.com/akutz/memconn"
	"github.com/atotto/webrtc-sdp-exchanger/apis"
	"github.com/atotto/webrtc-sdp-exchanger/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("failed to create firebase app: %s", err)
	}
	fsClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("failed to create firestore: %s", err)
	}
	defer fsClient.Close()

	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	svc := service.NewExchangeService(fsClient)
	apis.RegisterExchangeServiceServer(grpcServer, svc)

	reflection.Register(grpcServer)

	listener, err := memconn.Listen("memu", "grpc")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go func() {
		if err := grpcServer.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	conn, err := grpc.Dial(
		"grpc",
		grpc.WithInsecure(),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return memconn.DialContext(ctx, "memu", addr)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	mux := runtime.NewServeMux()
	apis.RegisterExchangeServiceHandler(ctx, mux, conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedOrigins: []string{
			"https://*",
		},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(mux)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}

	log.Printf("Listening on port %s", port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failure: %s", err)
	}
	grpcServer.GracefulStop()
}
