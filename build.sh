#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

echo "Go version:"
go version

echo "Tidying modules..."
go mod tidy

echo "Preparing output directory..."
OUTPUT_DIR="output"
mkdir -p "${OUTPUT_DIR}"

echo "Building library..."
go build -v ./...

echo "Build finished. Output directory: ${OUTPUT_DIR}"
