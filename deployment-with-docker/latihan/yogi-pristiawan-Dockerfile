FROM golang:1.18.5-alpine as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o binary cmd/api/main.go

FROM alpine:3.16.0
RUN apk --no-cache add bash
WORKDIR /app
COPY --from=build /app ./
EXPOSE 1234
ENTRYPOINT ["/app/binary"]