## Build
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /gofiberapp

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /gofiberapp /gofiberapp
COPY --from=build /app/service-account.json /service-account.json

EXPOSE 3000

ENTRYPOINT ["/gofiberapp"]