# example-service
A simple example microservice in Go.

# building & running
Build protobuf:
```shell script
~/projects/example-service$ protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative proto/service.proto
```

Quickly run with `go run`:
```shell script
~/projects/example-service$ go run .
```
Build and run with Docker:
```shell script
~/projects/example-service$ docker build -t example-service .
~/projects/example-service$ docker run -d -P example-service:latest
```

# todo
- Add promtheus metrics for gRPC
- Pull common functionality into a separate module