# Architecture of Volex DNS

Volex DNS is designed for high performance and privacy by combining concurrent upstream querying with DNS-over-TLS (DoT).

## High-Level Overview

The system consists of four primary components:

1.  **DNS Server (`internal/dnsserver`)**: Listens for incoming UDP DNS requests from local clients.
2.  **Memory Cache (`internal/cache`)**: Stores successful DNS responses to minimize latency for frequent queries.
3.  **Upstream Client (`internal/upstream`)**: Implements "Race Mode" querying over TLS.
4.  **Configuration Manager (`internal/config`)**: Handles JSON-based settings.

## Request Flow

1.  **Incoming Query**: A client sends a DNS query (UDP) to Volex DNS.
2.  **Special Handling**: The server checks if the query is a PTR request for `127.0.0.1`. If so, it immediately returns the `server_name` defined in the config.
3.  **Cache Lookup**: The server checks the in-memory cache using a key derived from the query name, type, and class.
    *   **Cache Hit**: The cached message is returned immediately.
4.  **Upstream Race**: On a cache miss, the server initiates a query to all configured DoT upstreams in parallel.
    *   Each upstream connection is established over TLS (port 853).
    *   The first successful response (excluding `ServerFailure`) is returned to the client.
5.  **Caching Results**: Successful responses are stored in the cache with a TTL derived from the response records (bounded by `cache_min_ttl_sec` and `cache_max_ttl_sec`).
6.  **Cleanup**: A background loop runs every 10 minutes to remove expired entries from the cache.

## Race Mode Logic

The Upstream Client uses Go's concurrency primitives (`goroutines` and `channels`):

```go
// Simplified logic from internal/upstream/client.go
for _, upstream := range uc.Upstreams {
    go func(server string) {
        resp, _, err := c.Exchange(r, server)
        if err == nil {
            respChan <- resp
        }
    }(upstream)
}

select {
case resp := <-respChan:
    return resp, nil
case <-time.After(uc.Timeout):
    return nil, errors.New("timeout")
}
```

This ensures that even if one upstream is slow or down, the client receives the fastest possible answer from the remaining pool.

## Security

*   **DNS-over-TLS (DoT)**: All external communication is encrypted using TLS, preventing eavesdropping and tampering by ISPs or attackers.
*   **EDNS0**: Supports Extension Mechanisms for DNS, including DNSSEC OK (DO) bit forwarding.
