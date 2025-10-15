FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .   

RUN go build -o /app/main ./cmd/main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.18

COPY --from=builder /app/main /main
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/image ./image

COPY .env .env

EXPOSE 8082

CMD ["sh", "-c", "echo $DB_URL && goose -dir ./migrations postgres \"$DB_URL\" up && ./main"]
