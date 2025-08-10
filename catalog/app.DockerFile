# Build stage
FROM golang:1.22-alpine3.20 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /src
COPY go.mod go.sum ./
COPY vendor vendor
COPY account account
RUN go build -mod vendor -o /app ./catalog/cmd/catalog

# Final minimal stage
FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /app .
EXPOSE 8080
CMD ["./app"]
