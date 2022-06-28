FROM golang:1.18-alpine AS builder

# Directory to work.
WORKDIR /build

# Copy and download dependencies using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Set environment variables and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch 

# Copy binary and config file from /build to root folder of scratch container
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Export necessary port
EXPOSE 8000

# Run the server
ENTRYPOINT [ "/apiserver" ]