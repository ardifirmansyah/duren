FROM golang:1.14.2-alpine3.11 AS builder

WORKDIR /go/src/github.com/ardifirmansyah/duren

COPY . .

RUN go mod download
RUN go build -o engine

FROM alpine:latest 

RUN mkdir /app
WORKDIR /app

EXPOSE 8080

COPY --from=builder /go/src/github.com/ardifirmansyah/duren/engine /app
COPY --from=builder /go/src/github.com/ardifirmansyah/duren/files/docker/wait-for.sh wait-for.sh
RUN chmod +x wait-for.sh

CMD sh ./wait-for.sh postgres:5432 -- /app/engine