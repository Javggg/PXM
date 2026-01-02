#!/bin/bash

APP_NAME="pxm"

BUILD_DIR="build"
mkdir -p $BUILD_DIR

echo "Building $APP_NAME for all platforms..."

# Linux amd64
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-linux-amd64
if [ $? -ne 0 ]; then echo "Failed to build Linux amd64"; exit 1; fi

# Windows amd64
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-windows-amd64.exe
if [ $? -ne 0 ]; then echo "Failed to build Windows amd64"; exit 1; fi

# macOS (amd64)
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-mac-amd64
if [ $? -ne 0 ]; then echo "Failed to build macOS amd64"; exit 1; fi

# macOS (arm64)
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/${APP_NAME}-mac-arm64
if [ $? -ne 0 ]; then echo "Failed to build macOS arm64"; exit 1; fi

echo "Build completed. Binaries are in the '$BUILD_DIR' folder."
