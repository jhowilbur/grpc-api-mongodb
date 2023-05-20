package main

import (
	"github.com/jhowilbur/grpc-api-mongodb/client/resources"
	pb "github.com/jhowilbur/grpc-api-mongodb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := resources.CreateBlog(c)
	resources.ReadBlog(c, id)
	resources.ReadBlog(c, "aNonExistingID")
	resources.UpdateBlog(c, id)
	resources.ListBlog(c)
	resources.DeleteBlog(c, id)
}
