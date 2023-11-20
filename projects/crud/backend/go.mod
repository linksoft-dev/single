module github.com/linksoft-dev/single/crud

go 1.21

replace github.com/linksoft-dev/single/comps/go v0.0.1 => ./../../../comps/go

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/graphql-go/graphql v0.8.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0
	github.com/linksoft-dev/single/comps/go v0.0.1
	github.com/sirupsen/logrus v1.9.2
	google.golang.org/grpc v1.58.0
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gofiber/fiber/v2 v2.49.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/graphql-go/handler v0.2.3 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.49.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
