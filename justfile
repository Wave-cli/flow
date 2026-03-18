default:
    @just --list

build:
    go build -o bin/flow .

test:
    go test ./... -v -count=1

test-cover:
    go test ./... -cover -count=1

coverage:
    go test ./... -coverprofile=cover.out
    go tool cover -html=cover.out -o cover.html
    @echo "Coverage report: cover.html"

clean:
    rm -rf bin/ cover.out cover.html

lint:
    go vet ./...

fmt:
    gofmt -s -w .

# ─── local setup ───────────────────────────────────────────────────────

# install flow plugin to ~/.wave/plugins for local testing
setup:
    #!/usr/bin/env bash
    set -euo pipefail
    mkdir -p ~/.wave/plugins/flow/bin
    cp bin/flow ~/.wave/plugins/flow/bin/flow
    cp Waveplugin ~/.wave/plugins/flow/Waveplugin
    echo "Installed flow to ~/.wave/plugins/flow"
