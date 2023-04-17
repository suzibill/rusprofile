package main

import (
	"context"
	"log"

	pb "github.com/suzibill/rusprofile/api/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client stub.
	client := pb.NewRusProfileClient(conn)

	// Call the GetCompanyInfo method with the INN parameter.
	inn := "5609026406"
	req := &pb.CompanyRequest{Inn: inn}
	resp, err := client.GetCompanyInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get company info: %v", err)
	}

	// Print the response.
	log.Printf("Company Info: INN=%s, KPP=%s, Name=%s, CEO=%s", resp.Inn, resp.Kpp, resp.Name, resp.DirectorName)
}
