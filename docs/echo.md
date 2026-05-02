# ECHO

Returns the string passed to it. Useful for testing client–server round-trips.

## Syntax

```
ECHO message
```

## Arguments

| Argument  | Type   | Required | Description              |
|-----------|--------|----------|--------------------------|
| `message` | string | Yes      | The string to echo back. |

## Return value

The message exactly as provided.

## Errors

- `message is required` — when called without an argument.

## Examples

```
> ECHO "hello"
"hello"

> ECHO
(error) message is required
```
