FROM golang:1.19-alpine3.16 AS builder

LABEL maintainer="GGGoingDown <gggoingdown0401@gmail.com>"

WORKDIR /app 

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o apiserver main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/apiserver /app/apiserver

COPY app.env app.env

# interupt signal inside container
RUN apk add dumb-init

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD /app/apiserver

EXPOSE 8080