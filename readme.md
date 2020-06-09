# example-service
A simple example microservice in Go.

# building & running
Quickly run with `go run`:
```go
~/projects/example-service$ go run .
```
Build and run with Docker:
```go
~/projects/example-service$ docker build -t example-service .
~/projects/example-service$ docker run -d -P example-service:latest
```

# todo
- Add promtheus metrics for gRPC
- Pull common functionality into a separate module