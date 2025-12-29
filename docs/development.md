# Development Guide

This document provides information on how to build, test, and contribute to Volex DNS.

## Prerequisites

*   **Go**: 1.22 or higher.
*   **Permissions**: Administrator/Root privileges are required to run the server on port 53.

## Project Structure

*   `cmd/volex-dns/`: Entry point of the application.
*   `internal/dnsserver/`: Core DNS handling logic.
*   `internal/upstream/`: Concurrent TLS client implementation.
*   `internal/cache/`: TTL-aware in-memory cache.
*   `internal/config/`: Configuration loading.

## Building from Source

To build the binary:

```bash
go build -o volex-dns ./cmd/volex-dns
```

## Running Tests

To run the project tests (if available):

```bash
go test ./...
```

## Contribution Guidelines

1.  **Fork the repository**.
2.  **Create a feature branch** (`git checkout -b feature/amazing-feature`).
3.  **Commit your changes** (`git commit -m 'Add some amazing feature'`).
4.  **Push to the branch** (`git push origin feature/amazing-feature`).
5.  **Open a Pull Request**.

### Code Style

This project follows standard Go conventions (`gofmt`). Please ensure your code is formatted before submitting.
