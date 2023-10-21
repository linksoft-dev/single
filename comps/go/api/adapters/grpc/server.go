package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kissprojects/single/comps/go/api"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

var (
	ServerEndpoint = "localhost:8080"

	mux = runtime.NewServeMux()
)

func New(port string) Adapter {
	if port == "" {
		port = "8080"
	}
	SetPort(port)
	return Adapter{port: port, GeneratePB: true, RunHttpServer: true}
}

func SetPort(port string) {
	ServerEndpoint = "localhost:" + port
}

type AppI interface {
	api.App
	Register(grpcServer *grpc.Server, mux *runtime.ServeMux) error
	GetInterceptor() grpc.UnaryServerInterceptor
}

// Adapter struct to save apps that implement grpc adapters and set port that grpc server should run
type Adapter struct {
	port          string
	RunHttpServer bool
	apps          []AppI
}

func (g *Adapter) Add(app AppI) {
	g.apps = append(g.apps, app)
}

// Run method that implements Adapter interface
func (g Adapter) Run() error {
	var err error
	listen, err := net.Listen("tcp", g.getAddress())
	if err != nil {
		log.Warnf("failed to listen on %s: %v", g.port, err)
		return err
	}

	// add interceptors from all apps
	var interceptors []grpc.ServerOption
	for _, app := range g.apps {
		interceptor := app.GetInterceptor()
		if interceptor == nil {
			continue
		}
		interceptors = append(interceptors, grpc.UnaryInterceptor(interceptor))
	}

	grpcServer := grpc.NewServer(
		interceptors...,
	)

	// register all apps
	for _, app := range g.apps {
		err := app.Register(grpcServer, mux)
		if err != nil {
			return err
		}
	}

	reflection.Register(grpcServer)
	// start the server whatever anything
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve gRPC over %s: %v", g.port, err)
		}
	}()

	eg := errgroup.Group{}

	if g.RunHttpServer {
		eg.Go(func() error {
			go runWebServer()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	// test server connection
	for {
		r := g.testGrpcServerConnection()
		if r {
			log.Infof("GRPC server listening on %s\n", g.port)
			return nil
		}
	}
}

// getAddress return the address of the current grpc connection
func (g Adapter) getAddress() string {
	return fmt.Sprintf(":%s", g.port)
}
func (g Adapter) testGrpcServerConnection() bool {
	conn, err := grpc.DialContext(context.Background(),
		g.getAddress(),
		grpc.WithInsecure(),
		grpc.FailOnNonTempDialError(true), // fail immediately if can't connect
		grpc.WithBlock())
	if err != nil {
		return false
	}
	return conn.GetState() == connectivity.Ready
}
func (g Adapter) GetName() string {
	return "Grpc"
}

func (g Adapter) GetApps() []api.App {
	apps := []api.App{}
	for _, app := range g.apps {
		apps = append(apps, app)
	}
	return apps
}

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
