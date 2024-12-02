# Command Examples

Install from a file, unpack to current directory:

```sh
tarsync .../my-file.json
```

Install from a file, unpack to a specified root directory:

```sh
tarsync .../my-file.json --unpack-to=~/.local
```

Remove already-extracted files ("local cache")

```sh
tarsync clean abcd1234 acef2468
```

Purge existing download targets by hash:

```sh
tarsync purge abcd1234 acef2468
```

## Unpack cache

Once the archive is confirmed present, need to determine whether to unpack freshly, or use the currently cached versions of unpacked files.

Default may be to re-use files, but supply a `--force-unpack` flag to override. This can be useful in CI for affirming package content.
