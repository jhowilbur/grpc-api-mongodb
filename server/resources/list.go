package resources

import (
	"context"
	"fmt"
	pb "github.com/jhowilbur/grpc-api-mongodb/proto"
	"github.com/jhowilbur/grpc-api-mongodb/server/entities"
	"github.com/jhowilbur/grpc-api-mongodb/server/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) ListBlogs(_ *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	ctx := context.Background()
	cur, err := pkg.Collection.Find(ctx, primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &entities.BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		stream.Send(entities.DocumentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}
