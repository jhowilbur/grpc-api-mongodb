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
	"log"
)

func (*Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", in)

	data := entities.BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := pkg.Collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
