#!/usr/bin/env bash

go version || exit 9

# Linux
echo "Linux"                                                       && \
GOOS=linux GOARCH=amd64 go build -o go-andotp-linux-x86_64 main.go && \
GOOS=linux GOARCH=arm64 go build -o go-andotp-linux-arm64  main.go && \

# macOS
echo "macOS"                                                        && \
GOOS=darwin GOARCH=amd64 go build -o go-andotp-macos-x86_64 main.go && \
GOOS=darwin GOARCH=arm64 go build -o go-andotp-macos-arm64  main.go && \

# Windows
echo "Windows"                                                             && \
GOOS=windows GOARCH=amd64 go build -o go-andotp-windows-x86_64.exe main.go && \
GOOS=windows GOARCH=arm64 go build -o go-andotp-windows-arm64.exe  main.go && \

echo "✅ DONE"
