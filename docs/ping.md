# PING

Returns the server's liveness response. Useful for testing if the server is alive and measuring latency.

## Syntax

```
PING [message]
```

## Arguments

| Argument  | Type   | Required | Description                           |
|-----------|--------|----------|---------------------------------------|
| `message` | string | No       | If provided, echoes back this string. |

## Return value

- `PONG` when called with no arguments.
- The message echoed back when called with an argument.

## Examples

```
> PING
PONG

> PING hello
hello
```
