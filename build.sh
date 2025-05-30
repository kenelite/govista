#!/bin/bash

set -e

APP_NAME=govista
OUTPUT_DIR=build

PLATFORMS=(
  "darwin/amd64"
  "darwin/arm64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
)

mkdir -p $OUTPUT_DIR

echo "Building $APP_NAME for multiple platforms..."

for PLATFORM in "${PLATFORMS[@]}"; do
  IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
  OUTPUT_NAME=$APP_NAME-$GOOS-$GOARCH
  if [ "$GOOS" = "windows" ]; then
    OUTPUT_NAME+=".exe"
  fi

  echo "Building for $GOOS/$GOARCH..."
  env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=1 go build -o $OUTPUT_DIR/$OUTPUT_NAME .

done

echo "All builds completed in $OUTPUT_DIR/"
