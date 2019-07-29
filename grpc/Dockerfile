FROM golang:1.12.0-alpine as build

RUN apk update && apk add git 

WORKDIR /go/src/github.com/walker2/k8s-goapm-examples/grpc

ENV CGO_ENABLED=0
COPY protobuf ./protobuf
COPY *.go ./

RUN go get
RUN go build -o grpc


FROM alpine:3.8

RUN apk update && apk add --no-cache ca-certificates

COPY --from=build /go/src/github.com/walker2/k8s-goapm-examples/grpc .
ENTRYPOINT ["/grpc"]