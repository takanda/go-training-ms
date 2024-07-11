FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY app/go.mod ./
RUN go mod download

# Copy the source code
COPY app/code ./
