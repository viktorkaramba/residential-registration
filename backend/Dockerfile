FROM golang:1.21.5-alpine3.19 AS builder

COPY . /backend/
WORKDIR /backend/

#build app
RUN go mod download
RUN go build -o ./bin/api cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /backend/bin/api .
COPY --from=0 /backend/config config/

EXPOSE 80

CMD ["./api"]