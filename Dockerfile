# This image is a microservice in golang for the Degree chaincode
FROM golang:1.23 AS builder

# 👇 Buildx tự động cung cấp TARGETOS và TARGETARCH
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
WORKDIR /app
COPY . .

# build cho đúng platform được truyền vào
RUN go build -o chaincode .

# Stage 2: Create a lightweight runtime image
FROM alpine

WORKDIR /app
COPY --from=builder /app/chaincode .

RUN chmod +x ./chaincode
CMD ["./chaincode"]