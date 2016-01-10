# Microservice Template

A production-ready microservice template built with Go.

## Features

- Health checks (`/health`, `/ready`)
- Prometheus metrics (`/metrics`)
- Environment-based configuration
- Docker support
- Graceful shutdown
- Request logging middleware

## Quick Start

### Local Development
```bash
go run main.go
```

### Docker
```bash
docker build -t microservice-template .
docker run -p 8080:8080 microservice-template
```

## Endpoints

- `GET /` - Service info
- `GET /health` - Health check
- `GET /ready` - Readiness check  
- `GET /metrics` - Prometheus metrics

## Configuration

Environment variables:
- `PORT` - Server port (default: 8080)
- `SERVICE_NAME` - Service name (default: microservice-template)
- `VERSION` - Service version (default: 1.0.0)

## Docker Compose

```yaml
version: '3'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - SERVICE_NAME=my-api
      - VERSION=1.0.0
```

## Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: microservice-template
spec:
  replicas: 3
  selector:
    matchLabels:
      app: microservice-template
  template:
    metadata:
      labels:
        app: microservice-template
    spec:
      containers:
      - name: api
        image: microservice-template:latest
        ports:
        - containerPort: 8080
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
```