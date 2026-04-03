#!/bin/bash

APP_NAME="gmon"
OUTPUT_DIR="./build"

mkdir -p $OUTPUT_DIR

platforms=(
    "windows/amd64"
    "windows/arm64"
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

for platform in "${platforms[@]}"; do
    GOOS="${platform%/*}"
    GOARCH="${platform#*/}"

    output="$OUTPUT_DIR/${APP_NAME}_${GOOS}_${GOARCH}"

    if [ "$GOOS" = "windows" ]; then
        output+=".exe"
    fi

    echo "Building $GOOS/$GOARCH..."

    GOOS=$GOOS GOARCH=$GOARCH go build -o "$output" .

    if [ $? -eq 0 ]; then
        echo "✓ $output"
    else
        echo "✗ Failed: $GOOS/$GOARCH"
    fi
done

echo ""
echo "Done! Binaries in $OUTPUT_DIR/"