# EXISTS

Returns the number of the given keys that exist in the store. Expired keys are treated as non-existent.

If the same key is provided multiple times it is counted multiple times.

## Syntax

```
EXISTS key [key ...]
```

## Arguments

| Argument | Type   | Required | Description                  |
|----------|--------|----------|------------------------------|
| `key`    | string | Yes      | One or more keys to check.   |

## Return value

The number of keys that exist. Ranges from `0` (none exist) to the total number of key arguments supplied.

## Errors

- `wrong number of arguments for 'exists' command` — when called with no arguments.

## Examples

```
> SET a 1
OK

> SET b 2
OK

> EXISTS a
1

> EXISTS a b
2

> EXISTS a missing
1

> EXISTS missing
0

> EXISTS a a
2
```
