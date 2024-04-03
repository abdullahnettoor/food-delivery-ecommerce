# # Use the official Golang image as the base image
# FROM golang:1.21.6-alpine

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum to download dependencies
# COPY go.mod go.sum ./

# # Download and install dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . ./

# # Build the application
# RUN go build -o main cmd/main.go

# # Expose the port that the application will run on
# EXPOSE 3000

# # Command to run the application
# CMD ["./main"]
# RUN CGO_ENABLED=0 GOOS=linux

FROM golang:1.21.6-alpine3.18 AS build
WORKDIR /app
COPY ./ /app
RUN mkdir -p /app/build
RUN go mod download
RUN go build -v -o /app/build/api ./cmd/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/build/api /
COPY --from=build /app/internal/view /internal/view
COPY --from=build /app/.env /
EXPOSE 8989
ENTRYPOINT ["/api"]