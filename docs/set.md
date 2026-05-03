# SET

Sets `key` to hold a string value. If `key` already exists, its value is overwritten.

## Syntax

```
SET key value [NX | XX] [EX seconds | PX milliseconds | EXAT timestamp | PXAT timestamp | KEEPTTL] [GET]
```

## Arguments

| Argument          | Type      | Description                                                                 |
|-------------------|-----------|-----------------------------------------------------------------------------|
| `key`             | string    | The key to set.                                                             |
| `value`           | string    | The value to store.                                                         |
| `NX`              | flag      | Only set the key if it does **not** already exist.                          |
| `XX`              | flag      | Only set the key if it **already** exists.                                  |
| `EX seconds`      | integer   | Set expiry as a relative duration in seconds.                               |
| `PX milliseconds` | integer   | Set expiry as a relative duration in milliseconds.                          |
| `EXAT timestamp`  | integer   | Set expiry as an absolute Unix timestamp in seconds.                        |
| `PXAT timestamp`  | integer   | Set expiry as an absolute Unix timestamp in milliseconds.                   |
| `KEEPTTL`         | flag      | Keep the existing TTL when overwriting the value.                           |
| `GET`             | flag      | Return the old value before setting. Returns nil if the key did not exist.  |

**Mutual exclusivity:**
- `NX` and `XX` cannot be used together.
- `EX`, `PX`, `EXAT`, `PXAT`, and `KEEPTTL` are mutually exclusive — only one may be provided.

## Return value

- `OK` on success.
- The old value (or nil if the key didn't exist) when `GET` is specified.
- nil when `NX` or `XX` condition is not met.

## Errors

- `missing arguments` — fewer than two arguments provided.
- `syntax error` — unknown option, duplicate option, or conflicting options.
- `invalid expire time in 'set' command` — expiry value is not a valid positive integer.

## Examples

```
> SET name helix
OK

> SET session abc123 EX 3600
OK

> SET token xyz PX 5000
OK

> SET event log EXAT 9999999999
OK

> SET event log PXAT 9999999999000
OK
```

Set only if the key does not exist:

```
> SET name helix NX
nil

> SET newkey value NX
OK
```

Set only if the key already exists:

```
> SET name helix XX
OK

> SET ghost value XX
nil
```

Keep the existing TTL when updating the value:

```
> SET name helix EX 3600
OK

> SET name updated KEEPTTL
OK
```

Return the old value while setting a new one:

```
> SET name helix
OK

> SET name world GET
helix

> SET nonexistent value GET
nil
```
