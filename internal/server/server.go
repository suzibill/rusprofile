package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/suzibill/rusprofile/internal/proto/proto"
	"github.com/suzibill/rusprofile/internal/service/parser"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	portHTTP = ":8080"
	portGRPC = ":8081"
)

type server struct {
	pb.UnimplementedRusProfileServer
}

func (s *server) GetCompanyInfo(ctx context.Context, req *pb.CompanyRequest) (*pb.Company, error) {
	company, err := parser.GetCompanyInfo(req.Inn)
	res := &pb.Company{
		Name:         company.Name,
		Inn:          company.INN,
		Kpp:          company.KPP,
		DirectorName: company.CEOFullname,
	}
	return res, err
}

func StartServer() {
	s := grpc.NewServer()
	grpcServer := &server{}
	pb.RegisterRusProfileServer(s, grpcServer)

	lis, err := net.Listen("tcp", portGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	// Create a gRPC-gateway ServeMux
	gwmux := runtime.NewServeMux()

	// Register the gRPC server implementation with the ServeMux
	err = pb.RegisterRusProfileHandlerServer(context.Background(), gwmux, grpcServer)
	if err != nil {
		log.Fatalf("failed to register gRPC server implementation: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterRusProfileHandlerFromEndpoint(context.Background(), gwmux, lis.Addr().String(), opts)
	if err != nil {
		log.Fatalf("failed to register gRPC-gateway: %v", err)
	}

	// Start an HTTP server
	server := &http.Server{
		Addr:    portHTTP,
		Handler: gwmux,
	}
	//log.Printf("starting server on port %s", port)
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	// Обработчик Swagger UI
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("dist"))))
	mux.Handle("/rusprofile.swagger.json", http.FileServer(http.Dir("gen/swaggerui")))
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//go func() {
	//	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("failed to listen and serve: %v", err)
	//	}
	//}()

	// Wait for a signal to stop the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Stop the server
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown server: %v", err)
	}
	s.GracefulStop()

	log.Println("server has been stopped")

}
