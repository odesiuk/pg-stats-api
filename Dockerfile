FROM golang:alpine AS build

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./ .

RUN go build -a -o /app/pg-stats-api

## ----
FROM alpine AS runtime

COPY --from=build /app/pg-stats-api .

ENTRYPOINT ["./pg-stats-api"]
