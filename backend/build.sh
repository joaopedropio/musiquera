docker buildx create --use --name multiarch-builder
docker buildx inspect --bootstrap
docker buildx build --platform linux/arm64 -t joaopio/musiquera:arm64 .
