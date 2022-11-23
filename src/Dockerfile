FROM golang:1.18.3-buster as builder

WORKDIR /go/src/app


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o app ./cmd/main.go

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y ca-certificates apache2-utils && \
    rm -rf /var/cache/apt/*

WORKDIR /go/src/app

COPY --from=builder /go/src/app/ /go/src/app/
COPY --from=builder /wait /wait

CMD ["./app"]
