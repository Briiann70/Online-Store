# Use golang image
FROM golang:1.23.2-alpine

# Set working directory
WORKDIR /app

# Copy everything to working directory
COPY . .

# Install dependencies
RUN go mod tidy

# Build aplikasi Go
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["go", "run", "main.go"]
