# Command Examples

Install from a file, unpack to current directory:

```sh
tarsync .../my-file.json
```

Install from a file, unpack to a specified root directory:

```sh
tarsync .../my-file.json --unpack-to=~/.local
```

Purge existing download targets by hash:

```sh
tarsync purge abcd1234 acef2468
```

