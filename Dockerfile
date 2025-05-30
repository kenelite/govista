FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    golang \
    gcc-aarch64-linux-gnu \
    g++-aarch64-linux-gnu \
    libx11-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libgl1-mesa-dev xorg-dev

ENV GOOS=linux
ENV GOARCH=arm64
ENV CC=aarch64-linux-gnu-gcc
ENV CGO_ENABLED=1

WORKDIR /app
COPY . .
RUN go build -o govista-linux-arm64 ./cmd/govista
