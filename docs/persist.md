# PERSIST

Removes the existing timeout on `key`, so the key will no longer expire and
persists indefinitely.

## Syntax

```
PERSIST key
```

## Arguments

| Argument | Type   | Required | Description                          |
|----------|--------|----------|--------------------------------------|
| `key`    | string | Yes      | The key whose TTL should be removed. |

## Return value

| Value | Meaning                                                    |
|-------|------------------------------------------------------------|
| `1`   | The timeout was removed; the key now persists indefinitely |
| `0`   | The key does not exist, or exists but has no timeout       |

## Errors

- `wrong number of arguments for 'persist' command` — incorrect number of arguments.

## Examples

```
> SET session token EX 100
OK

> TTL session
99

> PERSIST session
1

> TTL session
-1

> PERSIST session
0

> PERSIST missing
0
```
