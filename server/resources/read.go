package resources

import (
	"context"
	"fmt"
	pb "github.com/jhowilbur/grpc-api-mongodb/proto"
	"github.com/jhowilbur/grpc-api-mongodb/server/entities"
	"github.com/jhowilbur/grpc-api-mongodb/server/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (*Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &entities.BlogItem{}
	filter := bson.M{"_id": oid}

	res := pkg.Collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return entities.DocumentToBlog(data), nil
}
