FROM golang:1.12.0-alpine as build

RUN apk update && apk add git 

WORKDIR /go/src/github.com/walker2/k8s-goapm-examples/mux-span

ENV CGO_ENABLED=0
COPY *.go ./

RUN go get
RUN go build -o mux-span


FROM alpine:3.8

RUN apk update && apk add --no-cache ca-certificates

COPY --from=build /go/src/github.com/walker2/k8s-goapm-examples/mux-span .
ENTRYPOINT ["/mux-span"]