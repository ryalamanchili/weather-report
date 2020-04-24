# This first stage is defined as build_base, and referenced by second build
FROM golang:1.14.1-alpine

RUN apk update && apk add --no-cache git gcc \
&& rm -rf /var/cache/apk/*

WORKDIR /weather

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on

ADD . /weather

# Build
RUN CGO_ENABLED=0 go build -o /weather

EXPOSE 8081

# Run the binary program produced by the build step
CMD ["./weather"]

