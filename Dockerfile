FROM golang:alpine AS build

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build -a -o pg-stats-api

## ----
FROM alpine AS runtime

COPY --from=build /app/pg-stats-api .

ENTRYPOINT ["./pg-stats-api"]
