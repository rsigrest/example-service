FROM grpc/go:latest AS protobuf
WORKDIR /usr/src/proto
COPY ./proto .
RUN protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative service.proto

FROM golang:1.14.4-alpine3.12 AS build
WORKDIR /usr/src/app
COPY . .
COPY --from=protobuf /usr/src/proto/service.pb.go /usr/src/app/proto/service.pb.go
RUN go get -d -v ./... \
 && go build -v -x -o /usr/src/app/bin/example-service .

FROM alpine:3.12
COPY --from=build /usr/src/app/bin/example-service .
RUN chmod +x example-service
CMD ./example-service
