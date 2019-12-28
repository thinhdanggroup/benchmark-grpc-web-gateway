# Template for grpc-gateway 

ARG alpine_version=3.11
FROM golang:1.12.14-alpine$alpine_version AS build

RUN apk add --update --no-cache git
WORKDIR /app
# ENV GOPATH=/app

# Copy all of the staged files (protos plus go source)
COPY . /app/

# Download the go dependencies.
# RUN go get ./...
# RUN go mod init .

WORKDIR /app
# Build the gateway
RUN go build -o grpc_gateway src/pkg/main/main.go

FROM alpine:$alpine_version
WORKDIR /app
COPY --from=build /app/grpc_gateway /app/
COPY --from=build /app/config.yaml /app/
COPY --from=build /app/src/gen/pb-go/ping.swagger.json /app/

EXPOSE 80
ENTRYPOINT ["/app/grpc_gateway"]
