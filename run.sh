
docker buildx create --use
docker buildx build --platform linux/amd64,linux/arm64 -t metafi/token-cc:0.1.3 --push .
docker push metafi/token-cc:0.1.3