package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

var (
	ServerEndpoint = "localhost:8080"

	mux = runtime.NewServeMux()
)

// runWebServer run webserver based on all protobuffers
func runWebServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := generateSwagger(); err != nil {
		log.WithError(err).Errorf("error while generate swagger")
	}
	addr := ":8081"
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Infof("Webserver listening on in prefix /api %s", addr)
	return http.ListenAndServe(addr, mux)
}

type Services interface {
	Register(s *grpc.Server, httpServer *runtime.ServeMux) error
	GetInterceptor() grpc.UnaryServerInterceptor
}

func StartGrpcServer(port string, services ...Services) error {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Warnf("failed to listen on %s: %v", port, err)
		return err
	}

	// add interceptors from all apps
	var interceptors []grpc.ServerOption
	for _, service := range services {
		interceptor := service.GetInterceptor()
		if interceptor == nil {
			continue
		}
		interceptors = append(interceptors, grpc.UnaryInterceptor(interceptor))
	}

	grpcServer := grpc.NewServer(
		interceptors...,
	)

	// register all services
	for _, service := range services {
		err := service.Register(grpcServer, mux)
		if err != nil {
			return err
		}
	}

	reflection.Register(grpcServer)
	// start the server whatever anything
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve gRPC over %s: %v", port, err)
		}
	}()

	eg := errgroup.Group{}

	eg.Go(func() error {
		go runWebServer()
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	log.Infof("GRPC server listening on %s\n", port)
	return nil
}
