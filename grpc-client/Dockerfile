ARG alpine_version=3.11
FROM golang:1.12.14-alpine$alpine_version AS build

RUN apk add --update --no-cache git
WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . /app/


WORKDIR /app
# Build the gateway

RUN go build -o grpc_client

FROM alpine:$alpine_version
WORKDIR /app
COPY --from=build /app/grpc_client /app/

CMD ["/app/grpc_client"]
