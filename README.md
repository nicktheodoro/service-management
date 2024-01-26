# Service Management API

## Overview

This API, built with the Gin framework and Gorm, provides a foundation for managing services within Golang applications. It's currently in its initial stage!

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
swag init --dir cmd/api --parseDependency --output docs
```

   
## Contributing

Contributions are welcome! Please feel free to open issues or pull requests to suggest improvements or add new features.

## License

This project is licensed under the [MIT License](./LICENSE).
