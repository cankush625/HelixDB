# EXPIRE

Sets a timeout on `key`. After the timeout elapses, the key is automatically deleted.

## Syntax

```
EXPIRE key seconds
```

## Arguments

| Argument  | Type    | Required | Description                                |
|-----------|---------|----------|--------------------------------------------|
| `key`     | string  | Yes      | The key to set the expiry on.              |
| `seconds` | integer | Yes      | Time to live in seconds. Must be positive. |

## Return value

- `1` — the timeout was set successfully.
- `0` — the key does not exist.

## Errors

- `wrong number of arguments` — when called with anything other than exactly two arguments.
- `invalid expire time in 'expire' command` — when `seconds` is not a valid positive integer.

## Examples

```
> SET greeting hello
OK

> EXPIRE greeting 10
1

> EXPIRE nonexistent 10
0

> EXPIRE greeting -1
error: invalid expire time in 'expire' command
```
