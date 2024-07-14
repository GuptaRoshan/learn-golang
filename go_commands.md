## Basic Commands in Go

1. **`go run`**

   - Compiles and runs the specified Go program.
   - Example: `go run main.go`

2. **`go build`**

   - Compiles the packages and dependencies, producing an executable binary.
   - Example: `go build` or `go build main.go`

3. **`go install`**

   - Compiles and installs the packages, placing the executable binary in `$GOPATH/bin`.
   - Example: `go install`

4. **`go test`**

   - Runs the tests in the specified package.
   - Example: `go test` or `go test ./...` for all packages

5. **`go clean`**

   - Removes object files and cached files.
   - Example: `go clean`

6. **`go fmt`**

   - Formats Go source code according to the official style.
   - Example: `go fmt` or `go fmt ./...` for all packages

7. **`go vet`**

   - Examines Go source code and reports suspicious constructs.
   - Example: `go vet` or `go vet ./...` for all packages

8. **`go get`**

   - Downloads and installs packages and dependencies.
   - Example: `go get github.com/some/package`

9. **`go list`**

   - Lists the packages.
   - Example: `go list` or `go list ./...` for all packages

10. **`go mod init`**

    - Initializes a new module in the current directory.
    - Example: `go mod init mymodule`

11. **`go mod tidy`**

    - Adds missing and removes unused modules.
    - Example: `go mod tidy`

12. **`go env`**

    - Prints Go environment information.
    - Example: `go env`

13. **`go doc`**

    - Shows documentation for a package or symbol.
    - Example: `go doc fmt` or `go doc fmt.Println`

14. **`go version`**

    - Prints the Go version.
    - Example: `go version`

15. **`go help`**
    - Provides help for Go commands.
    - Example: `go help` or `go help build` for help on a specific command
