# This image is a microservice in golang for the Degree chaincode
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

# Build for Linux
RUN GOOS=linux GOARCH=amd64 go build -o chaincode .

# Stage 2: Create a lightweight runtime image
FROM alpine

WORKDIR /app
COPY --from=builder /app/chaincode .

RUN chmod +x ./chaincode
CMD ["./chaincode"]