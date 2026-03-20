# wave-flow - Workflow Automation Plugin

wave-flow is a plugin for wave that lets you define and run project commands from your Wavefile.

## Features

### Run Project Commands
Define common commands once in your Wavefile and run them with `wave flow <command>`.

```bash
wave flow build   # Run build command
wave flow dev     # Start development server
wave flow test    # Run tests
```

### List Available Commands
See all commands defined in your project.

```bash
wave flow --list
```

### Environment Variables
Set environment variables for specific commands.

```toml
[flow.dev]
cmd = "npm run dev"
env = { PORT = "3000", NODE_ENV = "development" }
```

### Success/Failure Callbacks
Run follow-up commands based on the result.

```toml
[flow.test]
cmd = "npm test"
on_success = "echo Tests passed!"
on_fail = "echo Tests failed. Check output above."
```

## Version History

### v0.2.2 (Current)
- Simplified message when no command specified
- Better help output formatting

### v0.2.1
- Updated to use wave-core SDK improvements

### v0.2.0
- Show help by default when no command specified
- Added `-h/--help` flag for usage information

### v0.1.x
- Initial release with command execution
- Environment variable support
- Success/failure callbacks

## Roadmap

### Watch Mode (Planned)
Automatically re-run commands when files change.

```toml
[flow.dev]
cmd = "go run ."
watch = ["*.go", "*.mod"]
```

### Command Chaining (Planned)
Run multiple commands in sequence.

### Parallel Execution (Planned)
Run independent commands in parallel.
