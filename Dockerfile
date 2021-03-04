FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/build

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app .

# Start fresh from a smaller image
FROM alpine
RUN apk add ca-certificates

COPY --from=build_base /tmp/build/out/app /app

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced by `go build`
CMD ["/app"]