FROM golang:1.18 AS base

# ENV work_dir_golang=/go/src/private-eco-sticker-car-back-end

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apt-get update && apt-get install gcc

# Move to working directory
WORKDIR /go/src/private-eco-sticker-car-back-end

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
# RUN go build -a cmd/server/main.go
ARG mode
ENV envMode=$mode
RUN cd cmd/server/ && go build -o ../../main -tags ${envMode}

# Use project from scratch
FROM alpine:3 AS server_build
RUN apk add --no-cache ca-certificates

# Copy dir server from image base to image alpine
WORKDIR /root/app
COPY --from=base /go/src/private-eco-sticker-car-back-end .

# Export necessary port
EXPOSE 1323

# Command to run when starting the container
CMD ./main