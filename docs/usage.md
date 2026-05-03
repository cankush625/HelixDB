# Usage Guide

This guide walks you through running HelixDB and using it to store, retrieve, and expire data.

## Starting the server

Clone the repository and run:

```bash
git clone https://github.com/cankush625/HelixDB.git
cd HelixDB
go run .
```

Or build a binary:

```bash
go build -o helixdb .
./helixdb
```

The server starts on port `6378`.

## Connecting

HelixDB speaks RESP, so you can connect using `redis-cli`:

```bash
redis-cli -p 6378
```

Or with netcat for raw access:

```bash
nc localhost 6378
```

> A native HelixDB CLI is on the roadmap. For now, `redis-cli` is the recommended client.

## Testing the connection

Once connected, verify the server is up:

```
> PING
PONG
```

## Storing and retrieving data

Store a value with [SET](./set.md) and retrieve it with [GET](./get.md):

```
> SET name helix
OK

> GET name
helix
```

If you try to get a key that doesn't exist, you get nil back:

```
> GET nonexistent
nil
```

## Setting expiry on keys

You can store a key with a TTL so it automatically disappears after a given time.

Set a key that expires in 60 seconds:

```
> SET session abc123 EX 60
OK
```

Or in milliseconds:

```
> SET token xyz PX 5000
OK
```

You can also set expiry on a key that already exists using [EXPIRE](./expire.md):

```
> SET greeting hello
OK

> EXPIRE greeting 10
1
```

A return value of `1` means the TTL was set. `0` means the key didn't exist.

Once a key expires, GET returns nil:

```
> GET session
nil
```

## Deleting keys

Delete one or more keys with [DEL](./del.md):

```
> SET a 1
OK

> SET b 2
OK

> DEL a b
2
```

The return value is the number of keys actually deleted. Keys that don't exist are ignored:

```
> DEL nonexistent
0
```

## Finding keys by pattern

Use [KEYS](./keys.md) to list keys matching a pattern:

```
> SET user:1 alice
OK

> SET user:2 bob
OK

> KEYS user:*
user:1
user:2

> KEYS *
user:1
user:2
```

See the [KEYS](./keys.md) doc for supported pattern syntax.

> **Note:** `KEYS` scans the entire store. Use it for debugging, not in hot paths.
