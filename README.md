# Service Management API

## Overview

This example API, crafted with the Gin framework and Gorm, stands as a mature solution for effective service management within Golang applications. By embracing Go's conventions, the project structure adheres to best practices. The REST architecture is implemented, aligning with Golang standards to ensure clean semantics and maintainability.

## Installation

Before running the application, ensure you have the necessary dependencies installed. Follow the steps below:

1. **Install Go (Golang):**

   Download and install Go from [golang.org](https://golang.org/dl/). Verify the installation by running:

   ```bash
   go version
   ```

2. **Download Dependencies:**

   Run the following command to download the project dependencies:

   ```bash
   go mod download
   ```

## Run the application

```bash
go run .\cmd\server\main.go
```

## Run tests

```shell script
go test -v ./test/...
```

## Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --dir cmd/server --parseDependency --output docs
```

   
## Contributing

Contributions are welcome! Please feel free to open issues or pull requests to suggest improvements or add new features.

## License

This project is licensed under the [MIT License](./LICENSE).
