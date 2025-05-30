#!/bin/bash

set -e

APP_NAME="govista"
OUTPUT_DIR="build"
GO_VERSION="1.24.3"

PLATFORMS=(
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
  "darwin/amd64"
  "darwin/arm64"
)

mkdir -p $OUTPUT_DIR

for PLATFORM in "${PLATFORMS[@]}"; do
  IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
  OUTPUT_NAME="${APP_NAME}-${GOOS}-${GOARCH}"
  [ "$GOOS" = "windows" ] && OUTPUT_NAME="${OUTPUT_NAME}.exe"

  echo "🔧 Building $OUTPUT_NAME ..."

  docker run --rm -v "$PWD":/app -w /app \
    -e GOOS=$GOOS -e GOARCH=$GOARCH -e CGO_ENABLED=1 -e CC=clang CXX=clang++\
    golang:${GO_VERSION} \
    go build -buildvcs=false -o ${OUTPUT_DIR}/${OUTPUT_NAME} .
done

echo "✅ All builds complete. Output in ./${OUTPUT_DIR}/"
