package grpc

import (
	"github.com/waykiss/wkgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func New(port string) Adapter {
	return Adapter{port: port}
}

type appInterface interface {
	wkgo.App
	Register(grpcServer *grpc.Server)
}

// Adapter struct to save apps that implement grpc adapters and set port that grpc server should run
type Adapter struct {
	port string
	apps []appInterface
}

func (g *Adapter) Add(app appInterface) {
	g.apps = append(g.apps, app)
}

// Run method that implements wkgo.Adapter interface
func (g Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", g.port)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", g.port, err)
	}
	grpcServer := grpc.NewServer()
	for _, app := range g.apps {
		app.Register(grpcServer)
	}
	reflection.Register(grpcServer)
	log.Printf("GRPC server listening on %s\n", g.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC over %s: %v", g.port, err)
	}
}

func (g Adapter) GetApps() (r []wkgo.App) {
	for _, app := range g.apps {
		r = append(r, app)
	}
	return
}
