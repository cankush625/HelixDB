# KEYS

Returns all key names in the store that match the given glob-style pattern. Expired keys are excluded from the result.

> **Warning:** `KEYS` iterates over all keys in the store. Avoid using it in production against large datasets — it blocks the store scan for the duration. Use it for debugging or administrative purposes only.

## Syntax

```
KEYS pattern
```

## Arguments

| Argument  | Type   | Required | Description                          |
|-----------|--------|----------|--------------------------------------|
| `pattern` | string | Yes      | Glob-style pattern to match against. |

### Supported pattern syntax

| Pattern   | Matches                                      |
|-----------|----------------------------------------------|
| `*`       | Any sequence of characters                   |
| `?`       | Any single character                         |
| `[abc]`   | Any one character in the set                 |
| `[a-z]`   | Any one character in the range               |

> **Note:** Keys containing `/` are not fully supported with wildcard patterns.

## Return value

List of matching key names. Empty if no keys match.

## Errors

- `wrong number of arguments` — when called with anything other than exactly one argument.
- `invalid pattern` — when the pattern syntax is malformed (e.g., unclosed `[`).

## Examples

```
> SET user:1 "alice"
OK

> SET user:2 "bob"
OK

> SET session "xyz"
OK

> KEYS *
1) "user:1"
2) "user:2"
3) "session"

> KEYS user:*
1) "user:1"
2) "user:2"

> KEYS user:?
1) "user:1"
2) "user:2"

> KEYS nonexistent*
(empty array)
```
