# Build stage for the new plugin
FROM golang:1.21-alpine as plugin-build

# Install git for fetching dependencies
RUN apk add --no-cache --update git

# Set working directory inside the container
WORKDIR /app

# Copy plugin go files and go.mod to working directory
# Copy go files and go.mod to working directory
COPY *.go ./
COPY go.mod go.sum ./

# Fetch dependencies and build the plugin
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-template-plugin

# Final stage
FROM alpine:3.14

# Install utilities and zip for unpacking
RUN apk --no-cache add curl jq bash zip yq

# Copy the go-template zip and unpack it
# COPY go-template.zip /app/go-template.zip
# RUN unzip /app/go-template.zip -d /app && chmod +x /app/go-template


RUN mkdir -m 777 -p /app \
    && ccurl -s -L -o /app/go-template https://app.harness.io/public/shared/tools/go-template/release/v0.4.4/bin/linux/amd64/go-template

RUN chmod +x /app/go-template


# Move the go-template binary to /usr/local/bin to make it available in PATH
RUN mv /app/go-template /usr/local/bin/

# Copy the template test files
COPY test /app/test

# Create a directory for test results
RUN mkdir /app/test/results

# Copy the plugin binary from the build stage
COPY --from=plugin-build /app/go-template-plugin /usr/local/bin/go-template-plugin
RUN chmod +x /usr/local/bin/go-template-plugin

# Set the working directory
WORKDIR /app

# Default command to run the new plugin
CMD ["go-template-plugin"]
