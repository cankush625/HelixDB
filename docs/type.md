# TYPE

Returns the type of the value stored at a key.

## Syntax

```
TYPE key
```

## Return value

One of the following strings:

| Value     | Meaning                                  |
|-----------|------------------------------------------|
| `string`  | A plain string value                     |
| `integer` | A whole number (e.g. `42`, `-1`)         |
| `float`   | A decimal number (e.g. `3.14`)           |
| `boolean` | `true` or `false`                        |
| `json`    | A valid JSON object or array             |
| `none`    | Key does not exist or has expired        |

## Errors

- `wrong number of arguments for 'type' command` — incorrect number of arguments.

## Examples

```
> SET counter 42
OK

> TYPE counter
integer

> SET ratio 3.14
OK

> TYPE ratio
float

> SET active true
OK

> TYPE active
boolean

> SET payload {"user":"alice","age":30}
OK

> TYPE payload
json

> SET greeting hello
OK

> TYPE greeting
string

> TYPE missing
none
```
