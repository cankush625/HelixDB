# CONFIG

Read and modify server configuration at runtime without restarting.

## Syntax

```
CONFIG GET parameter
CONFIG SET parameter value
```

## Subcommands

### CONFIG GET

Returns the current value of a configuration parameter.

```
CONFIG GET parameter
```

### CONFIG SET

Updates a configuration parameter at runtime.

```
CONFIG SET parameter value
```

## Supported parameters

| Parameter               | Type    | Default | Description                                              |
|-------------------------|---------|---------|----------------------------------------------------------|
| `hz`                    | integer | `1`     | Number of active expiry cycles per second. Must be > 0.  |
| `active-expire-enabled` | yes/no  | `yes`   | Enable or disable the active expiry background job.      |
| `maxmemory`             | integer | `0`     | Max memory in bytes. `0` means unlimited.                |

## Return value

- `CONFIG GET` — a two-element list: `[parameter, value]`.
- `CONFIG SET` — `OK` on success.

## Errors

- `wrong number of arguments for 'config' command` — incorrect number of arguments.
- `unknown subcommand '<name>' for 'config' command` — unrecognised subcommand.
- `unknown config parameter '<name>'` — unrecognised parameter name on GET.
- `invalid value for config parameter '<name>'` — value is out of range or wrong type on SET.

## Examples

```
> CONFIG GET hz
hz
1

> CONFIG SET hz 10
OK

> CONFIG GET hz
hz
10

> CONFIG GET active-expire-enabled
active-expire-enabled
yes

> CONFIG SET active-expire-enabled no
OK

> CONFIG GET maxmemory
maxmemory
0

> CONFIG SET maxmemory 1073741824
OK
```
