# Stage 1: Build
FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


WORKDIR /app/cmd/api
RUN go build -o /main .

FROM alpine:latest

WORKDIR /root/


COPY --from=builder /main .

EXPOSE 8080

CMD ["./main"]
