docker buildx create --use --name multiarch-builder
docker buildx inspect --bootstrap
docker buildx build --platform linux/arm64 -t yourusername/yourimage:arm64 .
docker buildx build --platform linux/arm64 -t yourusername/yourimage:arm64 --push .
