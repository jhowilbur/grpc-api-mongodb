package resources

import (
	"context"
	pb "github.com/jhowilbur/grpc-api-mongodb/proto"
	"log"
)

func DeleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}
