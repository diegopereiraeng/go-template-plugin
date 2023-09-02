
# Build stage for the new plugin
FROM golang:1.21-alpine as plugin-build

# Install git for fetching dependencies
RUN apk add --no-cache --update git

# Set working directory for the plugin
WORKDIR /plugin

# Copy plugin go files and go.mod to working directory
COPY main.go plugin.go ./
COPY go.mod go.sum ./
# Initialize go mod
RUN go mod init go-template-plugin
# Fetch dependencies and build the plugin
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-template-plugin

# Final stage
FROM alpine:3.14

# Install utilities and zip for unpacking
RUN apk --no-cache add curl jq bash zip yq

# Copy the go-template zip and unpack it
COPY go-template.zip /app/go-template.zip
RUN unzip /app/go-template.zip -d /app && chmod +x /app/go-template

# Copy the template test files
COPY test /app/test

# Create a directory for test results
RUN mkdir /app/test/results

# Copy the plugin binary from the build stage
COPY --from=plugin-build /plugin/go-template-plugin /app/go-template-plugin
RUN chmod +x /app/go-template-plugin

# Set the working directory
WORKDIR /app

# Default command to run the new plugin
CMD ["/app/go-template-plugin"]
