FROM golang:alpine

WORKDIR /

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY ./migrations ./migrations
CMD ["sh", "-c", "goose -dir /migrations postgres \"postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable\" up"]