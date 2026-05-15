# HelixDB

---
### Helix is an in-memory data store that acts as a cache.

<p align="center">
    <img src="./assets/logo.png" width="150">
</p>

## Getting Started

See the [Usage Guide](./docs/usage.md) for installation, connecting, and a walkthrough of common operations.

## Installation

## Supported Commands

| Command                              | Description                                   |
|--------------------------------------|-----------------------------------------------|
| [PING](./docs/ping.md)               | Test server liveness                          |
| [ECHO](./docs/echo.md)               | Echo a message back                           |
| [GET](./docs/get.md)                 | Get the value of a key                        |
| [SET](./docs/set.md)                 | Set key to a string value with optional TTL   |
| [DEL](./docs/del.md)                 | Delete one or more keys                       |
| [EXPIRE](./docs/expire.md)           | Set a TTL on an existing key (in seconds)     |
| [KEYS](./docs/keys.md)               | Find keys matching a pattern                  |
| [CONFIG](./docs/config.md)           | Read and modify server configuration          |
| [TYPE](./docs/type.md)               | Return the type of the value stored at a key  |
| [EXISTS](./docs/exists.md)           | Check if one or more keys exist               |
| [TTL](./docs/ttl.md)                 | Get remaining TTL of a key in seconds         |
| [PTTL](./docs/pttl.md)              | Get remaining TTL of a key in milliseconds    |
