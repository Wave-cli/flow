# wave-flow

Development workflow automation plugin for [wave](https://github.com/wave-cli/wave-core).

## What it does

wave-flow reads command definitions from the `[flow]` section of your Wavefile and executes them. Each command is an inline map with a required `cmd` field and optional callbacks and environment variables.

## Wavefile example

```toml
[project]
name = "my-app"
version = "1.0.0"

[flow]
build = { cmd = "go build -o bin/app", on_success = "echo done", env = { GOOS = "linux" } }
clean = { cmd = "rm -rf bin/" }
dev   = { cmd = "go run .", watch = ["*.go", "*.mod"] }
test  = { cmd = "go test ./...", on_fail = "echo tests failed" }
```

## Usage

```bash
# Run a command
wave flow build
wave flow clean

# List available commands
wave flow --list
```

## Schema fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `cmd` | string | yes | Shell command to execute |
| `on_success` | string | no | Command to run after successful execution |
| `on_fail` | string | no | Command to run after failed execution |
| `env` | map | no | Additional environment variables |
| `watch` | string/array | no | File patterns to watch for changes |

## How it works

1. wave-core reads the `[flow]` section from the Wavefile
2. wave-core validates the section against the flow schema (Waveschema)
3. wave-core passes the section as JSON on stdin to the flow binary
4. flow parses the command, executes it via `sh -c`, and runs callbacks

Commands must be defined as **inline maps** under `[flow]`. Nested headers like `[flow.build]` are rejected by wave-core's rules engine.

## Install

```bash
wave install wave-cli/wave-flow
```

## Development

```bash
just build    # Build binary
just test     # Run tests
just coverage # Generate coverage report
```

## License

MIT
