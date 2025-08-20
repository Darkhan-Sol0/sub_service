FROM golang:alpine AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./cmd ./cmd/
COPY ./internal ./internal/
COPY ./docs ./docs/
COPY ./logger ./logger/
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app/main.go

FROM alpine
WORKDIR /
COPY --from=builder ./app ./
COPY ./config ./config/
CMD ["./app"]