# My Go App Template

A simple Go web application with automated building and deployment using GitHub Actions.

## Features

- RESTful API endpoints
- Health check endpoint
- Automated testing
- Multi-platform builds (Linux, Windows, macOS)
- Docker support
- Automatic releases

## Local Development

### Prerequisites
- Go 1.21 or later
- Git

### Running locally
```bash
# Clone the repository
git clone https://github.com/yourusername/my-go-app.git
cd my-go-app

# Download dependencies
go mod tidy

# Run the application
go run main.go

# Or build and run
go build -o my-go-app
./my-go-app
```

The application will start on port 8080 (or the port specified in the `PORT` environment variable).

### Available endpoints
- `GET /` - Home page
- `GET /health` - Health check
- `GET /api/info` - API information

### Running tests
```bash
go test -v ./...
```

## GitHub Actions

This project uses GitHub Actions for:
- Running tests on every push and pull request
- Building binaries for multiple platforms
- Creating releases when tags are pushed
- Building and pushing Docker images

### Creating a release
1. Create and push a tag:
```bash
git tag v1.0.0
git push origin v1.0.0
```

2. GitHub Actions will automatically:
    - Build binaries for all platforms
    - Create a GitHub release
    - Upload the binaries as release assets

## Docker

### Building locally
```bash
docker build -t my-go-app .
docker run -p 8080:8080 my-go-app
```

### Using pre-built image
```bash
docker run -p 8080:8080 yourusername/my-go-app:latest
```