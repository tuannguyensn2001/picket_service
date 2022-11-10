ARG APP_HTTP_ADDRESS
ARG APP_GRPC_ADDRESS

FROM golang:1.18-alpine3.16 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o main src/server/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/error.yml .
COPY --from=builder /app/.env .

CMD ["/app/main","server"]