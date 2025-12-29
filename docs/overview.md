# Documentation Overview

Welcome to the Volex DNS documentation. This guide provides a comprehensive look at the project's internal workings, setup, and contribution guidelines.

## Quick Links

- 🏛️ **[Architecture](architecture.md)** — Deep dive into the "Race Mode" logic and system design.
- ⚙️ **[Configuration](configuration.md)** — Complete reference for `config.json` parameters.
- 🛠️ **[Development](development.md)** — Information for developers and contributors.

## Project Goal

Volex DNS aims to be the fastest possible DNS forwarder for local networks. By querying multiple upstream providers simultaneously and using an intelligent caching mechanism, it ensures that your DNS resolution is always handled by the fastest available server.

## Key Concepts

### Race Mode
The core of Volex DNS. It doesn't just wait for one server; it asks many and takes the first answer. This mitigates latency spikes from any single provider.

### Privacy First
By using **DNS-over-TLS (DoT)**, all your DNS traffic is encrypted. Your ISP or any middleman cannot see which domains you are resolving.

### Efficiency
Written in Go, Volex DNS is extremely lightweight and can easily run on low-power devices like a Raspberry Pi or a home router.
