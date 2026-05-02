# GET

Returns the value stored at `key`. Returns nil if the key does not exist or has expired.

## Syntax

```
GET key
```

## Arguments

| Argument | Type   | Required | Description         |
|----------|--------|----------|---------------------|
| `key`    | string | Yes      | The key to look up. |

## Return value

- The value stored at `key`.
- nil if the key does not exist or has expired.

## Errors

- `wrong number of arguments` — when called with anything other than exactly one argument.

## Examples

```
> SET name helix
OK

> GET name
helix

> GET nonexistent
nil
```
