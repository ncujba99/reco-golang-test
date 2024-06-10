FROM golang:1.21

WORKDIR /app

# Copy all project files from context to working directory
COPY . .

# Install dependencies using Go modules
RUN go mod download

RUN go build -o goserver .

# # Final stage with minimal image
# FROM alpine:latestup

# # Copy the built binary from the builder stage
# COPY --from=builder /app/main /app

# # Set the working directory for the final container
# WORKDIR /app


CMD ["./main"]