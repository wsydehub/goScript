#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

echo "Go version:"
go version

echo "Tidying modules..."
go mod tidy

echo "Generating ANTLR sources..."
JAVA17_HOME=""
CANDIDATES=()
if [ -n "${JAVA17_HOME:-}" ]; then
  CANDIDATES+=("${JAVA17_HOME}")
fi
if [ -n "${JAVA_HOME_17:-}" ]; then
  CANDIDATES+=("${JAVA_HOME_17}")
fi
if [ -n "${JDK17_HOME:-}" ]; then
  CANDIDATES+=("${JDK17_HOME}")
fi
JAVA_HOME_FROM_TOOL=$(/usr/libexec/java_home -v 17 2>/dev/null || true)
if [ -n "${JAVA_HOME_FROM_TOOL}" ]; then
  CANDIDATES+=("${JAVA_HOME_FROM_TOOL}")
fi
if [ -d "/opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk/Contents/Home" ]; then
  CANDIDATES+=("/opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk/Contents/Home")
fi
if [ -d "/usr/local/opt/openjdk@17/libexec/openjdk.jdk/Contents/Home" ]; then
  CANDIDATES+=("/usr/local/opt/openjdk@17/libexec/openjdk.jdk/Contents/Home")
fi
for CANDIDATE in "${CANDIDATES[@]}"; do
  if [ -x "${CANDIDATE}/bin/java" ]; then
    JAVA17_VERSION=$("${CANDIDATE}/bin/java" -version 2>&1 | head -n 1 || true)
    if echo "${JAVA17_VERSION}" | grep -q '"17'; then
      JAVA17_HOME="${CANDIDATE}"
      break
    fi
  fi
done
if [ -z "${JAVA17_HOME}" ]; then
  echo "Java 17 not found. Set JAVA17_HOME to a JDK 17 installation."
  exit 1
fi
"${JAVA17_HOME}/bin/java" -jar antlr4-4.13.1-complete.jar -Dlanguage=Go -package goScript -visitor -listener GoScript.g4

echo "Preparing output directory..."
OUTPUT_DIR="output"
mkdir -p "${OUTPUT_DIR}"

echo "Building library..."
go build -v ./...

echo "Build finished. Output directory: ${OUTPUT_DIR}"
