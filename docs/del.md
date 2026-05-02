# DEL

Removes one or more keys. Keys that do not exist are silently ignored and not counted.

## Syntax

```
DEL key [key ...]
```

## Arguments

| Argument | Type   | Required | Description                      |
|----------|--------|----------|----------------------------------|
| `key`    | string | Yes      | One or more keys to delete.      |

## Return value

The number of keys that were actually deleted. Keys that did not exist are not counted.

## Errors

- `wrong number of arguments` — when called with no arguments.

## Examples

```
> SET a 1
OK

> SET b 2
OK

> DEL a b nonexistent
(integer) 2

> DEL nonexistent
(integer) 0
```
