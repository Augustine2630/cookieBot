FROM alpine:latest

# Install necessary dependencies for running Go applications
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY app app

RUN chmod +x /app/app

# Expose port 8080 (or the port your app uses)
EXPOSE 80

# Command to run the executable
CMD ["./app"]
