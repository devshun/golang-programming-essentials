FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod init aozora-search
RUN go mod tidy
RUN go build -o aozora-search

FROM alpine:latest

COPY --from=builder /app/aozora-search /aozora-search

ENTRYPOINT ["/aozora-search"]