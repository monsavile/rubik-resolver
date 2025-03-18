FROM golang:1.24.1-alpine AS builder
COPY . /go/src/rubik-resolver
WORKDIR /go/src/rubik-resolver
RUN go mod download
RUN go build -o ./bin/grpc_server cmd/grpc_server/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/rubik-resolver/bin/grpc_server .
COPY .env .
CMD ["./grpc_server"]
