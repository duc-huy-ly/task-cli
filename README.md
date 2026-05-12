# task-cli


A small command-line task manager written in Go. Provides basic task operations (add, remove, update, mark in-progress/done, and list) implemented in the `core` package with a thin `main` CLI.
https://roadmap.sh/projects/task-tracker

## Requirements

- Go 1.20+ (module-aware). The project uses a Go module: `github.com/dly/task-cli`.

## Build

From the repository root:

```bash
go build
# or run directly
go run .
```

## Usage
You can choose to run it with go or with the compiled executable.
```bash 
# Running with the executable after building
./task-cli <command> <optional-args>
./task-cli -h # show help
./task-cli list 
./task-cli add "Go to the dentist"
./task-cli rm 1 # removes the task of id 1
```
The CLI recognizes simple commands (see `main.go`). Example run from project root:

```bash
# run the program
go run . add "Make cookies"
go run . list
go run . -h   # show help
```

## Tests

Run unit tests for the project or specific packages:

```bash
go test ./... -v
go test ./core -run TestId -v
```

## Design notes


- IDs are generated sequentially by the in-memory generator.
- Exported names (Uppercase) are required to use `core` from `main` or other packages.

## Contributing

Issues, bug fixes, and tests are welcome. Keep changes small and include tests for new behavior.
