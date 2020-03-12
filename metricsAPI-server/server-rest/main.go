package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	metricsAPIpb "k8s-api/protofiles"
)

func StartHTTP(grpcPort string, httpPort string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// register gateway handler for gRPC server endpoint.
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	addr := grpcPort
	if err := metricsAPIpb.RegisterClientAPIServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		fmt.Println("failed to register gRPC server endpoint= ", err)
	}
	fmt.Println("registered the gateway handler for endpoint:", grpcPort)
	gw := &http.Server{Addr: httpPort, Handler: mux}

	fmt.Println("starting HTTP/REST gateway on port:", httpPort)

	//start HTTP server (and proxy calls to gRPC server endpoint).
	err := gw.ListenAndServe()
	if err != nil {
		fmt.Println("Error in HTTP is:")
	}

}

func main() {
	StartHTTP(":9090", ":50000")
}
