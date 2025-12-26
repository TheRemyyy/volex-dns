<div align="center">

# Volex DNS

**The Fastest DNS Server in the Local Universe**

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/TheRemyyy/volex-dns/build.yml?style=flat-square)](https://github.com/TheRemyyy/volex-dns/actions)

*A high-performance, caching DoT (DNS over TLS) forwarder customized for speed and privacy.*

[Features](#features) • [Installation](#installation) • [Configuration](#configuration)

</div>

---

## Overview

Volex DNS is a lightweight, concurrent DNS forwarder written in Go. It connects to multiple upstream DNS-over-TLS (DoT) providers simultaneously and returns the fastest response, ensuring your browsing experience is always snappy and secure.

### Key Features

- **⚡ Race Mode** — Queries multiple upstreams (Cloudflare, Google, Quad9) in parallel; first response wins.
- **🔒 Secure** — All upstream connections use DNS over TLS (DoT) on port 853.
- **🧠 Smart Caching** — In-memory caching with configurable TTLs to reduce latency for frequent queries.
- **🛠️ Self-Service** — Responds to local PTR requests for instant identifying.
- **🚀 Efficiency** — Optimized for low memory and CPU usage.

## <a id="installation"></a>📦 Installation

### From Source

```bash
git clone https://github.com/TheRemyyy/volex-dns.git
cd volex-dns
go build -o volex-dns ./cmd/volex-dns
```

## <a id="configuration"></a>⚙️ Configuration

The server is configured via `config.json`.

```json
{
  "upstreams": [
    "1.1.1.1:853",
    "8.8.8.8:853",
    "9.9.9.9:853"
  ],
  "server_name": "volex.dns",
  "bind_addr": ":53",
  "upstream_timeout_ms": 2000,
  "cache_max_ttl_sec": 3600,
  "cache_min_ttl_sec": 60
}
```

## Running

Run the binary with Administrator privileges (required for binding to port 53):

```bash
sudo ./volex-dns
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
<sub>Built with ❤️ and Go</sub>
</div>
