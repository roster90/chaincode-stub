# This image is a microservice in golang for the Degree chaincode
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

# ✅ Build static binary để chạy trong slim image
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chaincode .

# Runtime stage
FROM debian:bullseye-slim

COPY --from=builder /app/chaincode /app/chaincode
WORKDIR /app

RUN chmod +x ./chaincode

CMD ["./chaincode"]