FROM golang:1.12.0-alpine as build

RUN apk update && apk add git 

WORKDIR /go/src/walker2/go-apm-test

ENV CGO_ENABLED=0
COPY *.go ./

RUN go get
RUN go build


FROM alpine:3.8

RUN apk update && apk add --no-cache ca-certificates

COPY --from=build /go/src/walker2/go-apm-test .
ENTRYPOINT ["/go-apm-test"]