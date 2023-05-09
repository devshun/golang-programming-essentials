FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod init aozora-collector
RUN go mod tidy
RUN go build -o aozora-collector

FROM alpine:latest

COPY --from=builder /app/aozora-collector /aozora-collector

ENTRYPOINT ["/aozora-collector"]