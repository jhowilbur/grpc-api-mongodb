# Go Blog Project with gRPC and MongoDB

## Description


### Installation

1. Generate the gRPC code from the proto file
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/blog.proto
```

2. Run the docker-compose file
```bash
docker-compose up
```

3. Run the server
```bash
go run server/server.go
```

