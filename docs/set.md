# SET

Sets `key` to hold a string `value`. If `key` already exists, its value is overwritten regardless of type.

## Syntax

```
SET key value [EX seconds | PX milliseconds]
```

## Arguments

| Argument         | Type    | Required | Description                                          |
|------------------|---------|----------|------------------------------------------------------|
| `key`            | string  | Yes      | The key to set.                                      |
| `value`          | string  | Yes      | The value to store.                                  |
| `EX seconds`     | integer | No       | Set expiry in seconds. Must be a positive integer.   |
| `PX milliseconds`| integer | No       | Set expiry in milliseconds. Must be a positive integer. |

`EX` and `PX` are mutually exclusive — providing both is a syntax error.

## Return value

`OK` on success.

## Errors

- `missing arguments` — fewer than two arguments provided.
- `syntax error` — unknown option, duplicate option, or both `EX` and `PX` provided.
- `invalid expire time in 'set' command` — expiry value is not a valid positive integer.

## Examples

```
> SET name "helix"
OK

> SET counter 42
OK

> SET session "abc123" EX 3600
OK

> SET token "xyz" PX 5000
OK

> SET key value EX 10 PX 10000
(error) syntax error
```
