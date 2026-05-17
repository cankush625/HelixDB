# TTL

Returns the remaining time-to-live of a key in seconds.

## Syntax

```
TTL key
```

## Return value

| Value | Meaning                             |
|-------|-------------------------------------|
| `N`   | Remaining TTL in seconds (N ≥ 1)    |
| `-1`  | Key exists but has no TTL           |
| `-2`  | Key does not exist or has expired   |

## Errors

- `wrong number of arguments for 'ttl' command` — incorrect number of arguments.

## Examples

```
> SET counter 42 EX 100
OK

> TTL counter
99

> SET greeting hello
OK

> TTL greeting
-1

> TTL missing
-2
```
