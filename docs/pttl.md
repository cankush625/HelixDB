# PTTL

Returns the remaining time-to-live of a key in milliseconds.

## Syntax

```
PTTL key
```

## Return value

| Value | Meaning                                  |
|-------|------------------------------------------|
| `N`   | Remaining TTL in milliseconds (N ≥ 1)   |
| `-1`  | Key exists but has no TTL               |
| `-2`  | Key does not exist or has expired        |

## Errors

- `wrong number of arguments for 'pttl' command` — incorrect number of arguments.

## Examples

```
> SET counter 42 EX 100
OK

> PTTL counter
99743

> SET greeting hello
OK

> PTTL greeting
-1

> PTTL missing
-2
```
