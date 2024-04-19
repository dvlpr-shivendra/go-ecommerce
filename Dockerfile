FROM golang:1.22-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN go build -o /ecommerce

EXPOSE 9090

CMD ["/ecommerce"]

