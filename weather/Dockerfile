# This first stage is defined as build_base, and referenced by second build
FROM golang:1.14.1-alpine AS build_base

RUN apk update && apk add --no-cache git gcc \
&& rm -rf /var/cache/apk/*

WORKDIR /weather

COPY go.mod .
RUN GO111MODULE=0n

ADD . /weather

# Build
RUN CGO_ENABLED=0 go build -o /weather

# Final stage
COPY --from=build_base /go/bin/weather /go/bin/weather

EXPOSE 9090

# Run the binary program produced by the build step
CMD ["./weather"]

