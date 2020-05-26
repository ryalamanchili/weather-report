# This first stage is defined as build_base, and referenced by second build
FROM golang:1.14.1-alpine as build_base

RUN apk update && apk add --no-cache git gcc \
&& rm -rf /var/cache/apk/*

WORKDIR /weather

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on

ADD . /weather

# Build
RUN CGO_ENABLED=0 go build -o /go/bin/weather .

# Final stage
FROM alpine

COPY --from=build_base /go/bin/weather .

ENTRYPOINT ["./weather"]

EXPOSE 8080

