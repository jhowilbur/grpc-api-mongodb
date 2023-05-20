package resources

import (
	"context"
	pb "github.com/jhowilbur/grpc-api-mongodb/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func ListBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
