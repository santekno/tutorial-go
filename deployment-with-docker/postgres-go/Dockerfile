# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/album-api main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/.env /app/.env
COPY --from=build /app/album-api /app/album-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/album-api"]