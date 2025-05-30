#!/bin/bash
set -euo pipefail

echo "Start cross-platform building govista..."

export CGO_ENABLED=0
export CC=""
export CXX=""

VERSION=$(git describe --tags --always --dirty || echo "v0.0.0-dev")
echo "Building version: $VERSION"

OUTPUT_DIR=build
mkdir -p "$OUTPUT_DIR"

platforms=(
  "linux amd64"
  "windows amd64"
  "darwin amd64"
  "darwin arm64"
)

for platform in "${platforms[@]}"; do
  set -- $platform
  GOOS=$1
  GOARCH=$2

  if [[ "$GOOS" == "windows" ]]; then
    CGO_ENABLED=0
    CC=x86_64-w64-mingw32-gcc
    CXX=""
    EXT=".exe"
  else
    CGO_ENABLED=1
    CC=clang
    CXX=clang++
    EXT=""
  fi

  echo "Building for $GOOS/$GOARCH (CGO_ENABLED=$CGO_ENABLED)..."

  OUTPUT_NAME="govista-${GOOS}-${GOARCH}${EXT}"

  env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=$CGO_ENABLED CC=$CC CXX=$CXX \
    go build -ldflags "-X main.version=$VERSION" -o "$OUTPUT_DIR/$OUTPUT_NAME" main.go

  echo "Built $OUTPUT_NAME"
done

echo "All builds finished in $OUTPUT_DIR"
