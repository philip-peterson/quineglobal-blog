# QUINE Global Blog

A small Go web application serving the QUINE Global blog at [blog.quineglobal.com](https://blog.quineglobal.com).

Built with [gomponents](https://www.gomponents.com/) and HTMX.

## Running locally

### Prerequisites

- Go 1.23 or newer

### Start the development server

```bash
make start
```

Or directly:

```bash
go run ./cmd/app
```

The site will be available at **http://localhost:8080**.

No database or external services are required.

### Other commands

| Command             | Description                     |
|---------------------|---------------------------------|
| `make start`        | Run the development server      |
| `make test`         | Run tests                       |
| `make lint`         | Run golangci-lint               |
| `make build-docker` | Build a production Docker image |

## Deploying

The [CD workflow](.github/workflows/cd.yml) automatically builds a multi-platform Docker image and pushes it to the GitHub container registry (GHCR), tagged with the commit hash as well as `latest`.

You can build the image locally with:

```shell
make build-docker
```

> Note: You need the containerd image store enabled in Docker Desktop for multi-platform builds.