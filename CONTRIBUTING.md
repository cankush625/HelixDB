# Contributing to HelixDB

Thanks for taking the time to contribute!

## Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- Git

## Setup

1. Fork the repository on GitHub.

2. Clone your fork:
   ```bash
   git clone https://github.com/<your-username>/HelixDB.git
   cd HelixDB
   ```

3. Build and run:
   ```bash
   go run .
   ```

4. Verify it's working:
   ```bash
   redis-cli -p 6378 PING
   # PONG
   ```

## Project structure

```
cmd/        # Command implementations (one file per command + test file)
common/     # Shared types, errors, constants, and RESP helpers
db/         # In-memory store and active expiry background job
exc/        # Command executor — routes parsed commands to cmd/ handlers
resp/       # RESP protocol: parser, request handler, reply
docs/       # User-facing documentation
main.go     # Entry point — starts the TCP server
```

## Making changes

Create a branch from `master` using the format `<github-username>/<short-description>`:

```bash
git checkout -b ankushchavan/expire-command
```

Keep each branch focused on a single feature or fix.

## Commit messages

We use [Conventional Commits](https://www.conventionalcommits.org/), scoped to the area being changed:

```
type(scope): short description (#issue-number)
```

**Types:**

| Type       | When to use                              |
|------------|------------------------------------------|
| `feat`     | Adding a new feature                     |
| `fix`      | Fixing a bug                             |
| `refactor` | Code change that isn't a feature or fix  |
| `docs`     | Documentation only                       |
| `test`     | Adding or updating tests                 |
| `chore`    | Maintenance tasks, dependency updates, CI/CD changes |

**Examples from this repo:**

```
feat(cmd): implement EXPIRE command (#53)
fix(set): validate conflicting EX and PX options (#49)
refactor(cmd): replace []string with Cmd struct (#55)
docs: add command-specific documentation (#57)
test(cmd): add unit tests for DEL command (#54)
chore(ci): remove snyk workflow from actions (#36)
```

## Tests

Run the full test suite before opening a PR:

```bash
go test ./...
```

Every new command in `cmd/` should have a corresponding `_test.go` file. Follow the pattern of existing test files like `cmd/set_test.go` or `cmd/expire_test.go`.

## Opening a pull request

- Link the PR to its GitHub issue in the description.
- Keep the PR focused — one feature or fix per PR.
- Make sure `go test ./...` passes.
- Use the same commit message format for the PR title.

## Reporting bugs and requesting features

Open an issue on GitHub. For bugs, include steps to reproduce. For features, describe the use case.
