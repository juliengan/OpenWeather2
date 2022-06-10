FROM golang:1.17.11-alpine3.16 AS build-env

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk add --update --no-cache ca-certificates git

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN go build -o /efrei-devops-tp2

RUN apk update && apk add \
curl \
vim
EXPOSE 8081