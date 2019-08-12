FROM golang:alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

LABEL maintainer="github@shanaakh.pro"

RUN apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/server ./cmd/server/main.go

FROM alpine

RUN apk update && apk upgrade && apk add curl

WORKDIR /app

COPY --from=build /go/bin/server /server
COPY ./config /config

EXPOSE 8080

HEALTHCHECK --interval=1m --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1

ENTRYPOINT ["./server"]
