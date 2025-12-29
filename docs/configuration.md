# Configuration Guide

Volex DNS is configured via a single `config.json` file located in the project root.

## Configuration Options

| Parameter | Type | Default (Example) | Description |
| :--- | :--- | :--- | :--- |
| `upstreams` | Array of strings | `["1.1.1.1:853", "8.8.8.8:853"]` | List of DNS-over-TLS servers to query. Must include port (usually 853). |
| `server_name` | String | `"volex.dns"` | Hostname used for local identification and PTR responses. |
| `bind_addr` | String | `":53"` | Address and port to listen on for incoming DNS queries. |
| `upstream_timeout_ms` | Integer | `2000` | Timeout in milliseconds for upstream queries before giving up. |
| `cache_max_ttl_sec` | Integer | `3600` | Maximum time (in seconds) a record can stay in cache, regardless of its original TTL. |
| `cache_min_ttl_sec` | Integer | `60` | Minimum time (in seconds) a record will stay in cache, even if its original TTL is lower. |

## Example `config.json`

```json
{
  "upstreams": [
    "1.1.1.1:853",
    "1.0.0.1:853",
    "8.8.8.8:853",
    "8.8.4.4:853",
    "9.9.9.9:853"
  ],
  "server_name": "volex.dns",
  "bind_addr": ":53",
  "upstream_timeout_ms": 1500,
  "cache_max_ttl_sec": 86400,
  "cache_min_ttl_sec": 300
}
```

## Recommended Upstreams

For best performance and privacy, consider these providers:

*   **Cloudflare**: `1.1.1.1:853`, `1.0.0.1:853`
*   **Google**: `8.8.8.8:853`, `8.8.4.4:853`
*   **Quad9**: `9.9.9.9:853`, `149.112.112.112:853`
*   **AdGuard DNS**: `94.140.14.14:853`, `94.140.15.15:853` (includes ad-blocking)
