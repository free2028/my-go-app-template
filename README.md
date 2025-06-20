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

## Configuration

The application can be configured using environment variables:
- `PORT` - Port to listen on (default: 8080)
```

## 使用说明

1. **创建仓库**：在GitHub上创建一个新仓库
2. **克隆并添加代码**：将上述文件添加到你的仓库中
3. **配置secrets**（如果需要Docker功能）：
   - `DOCKER_USERNAME` - Docker Hub用户名
   - `DOCKER_PASSWORD` - Docker Hub访问令牌
4. **推送代码**：
   ```bash
   git add .
   git commit -m "Initial commit"
   git push origin main
```
5. **创建发布版本**：
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

这个配置会自动：
- 运行测试
- 构建多平台二进制文件
- 创建GitHub释出
- 构建Docker镜像（可选）

每次推送代码或创建tag时，GitHub Actions都会自动运行构建流程。