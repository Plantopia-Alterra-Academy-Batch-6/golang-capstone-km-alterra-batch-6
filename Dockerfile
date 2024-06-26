#Build stage
FROM golang:1.22-alpine3.19 AS build-stage

WORKDIR /app

COPY . .

RUN apk add --no-cache tzdata

RUN go mod download

RUN go build -o /goapp

#Release stage
FROM alpine:3.19 AS build-release-stage

WORKDIR /

# Copy tzdata from build stage to release stage
COPY --from=build-stage /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-stage /goapp /goapp
COPY --from=build-stage /app/modules /app/modules

EXPOSE 8080

ENTRYPOINT ["./goapp"]
