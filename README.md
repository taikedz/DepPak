# TarSync

A small utility inspired by the "dependencies" section in `build.zig.zon`.

Allows setting dependencies from a file pointing to tarballs, including hash-checking.

Tarballs are downloaded to `~/.local/var/tarsync/tarballs/`

Reads a `tarsync.json` file which has the following format:

```js
[
    {
        "hash": "abcd1234",
        "url": "https://some_url/file.tgz",
        "deploy" : {
            // copy from archive's `appcode/` location to the deploy location's `src/feature` subdir
            "appcode/" : ["src/feature"]
            // The same source location can be copied to any number of destination locations
        }
    },
    {
        // ... other tarball specs ...
    }
]
```

* This will store a tarball at `~/.local/var/tarsync/tarballs/abcd1234/file.tgz`
* TarSync will unpack the tarball, and move the contents of `appcode/` into `./src/feature`

This syncs all the specified tarballs, validates the hash, and unpacks it to a location.

Hash can be specified as `"-"` to cause the downloaded item's hash to be printed to console. Unpack will not proceed in this case. Re-run with the hash populated to the file to proceed.

See [command examples](command_examples.md) for additional information

## Applications

This can be used where any bunch of distributable files is to be collected.

I would have happily used this for distributing a spec list of Minetest mods for a server, for example.

Any language, project or application that doesn't have its own bundle distribution system can use this.

As a libaray, this can be integrated into an application for it to provide extensibility with user plugins.

## Goals

* Simple specification of file
* Standalone commandline utility as single binary with no runtime dependencies
* Available as a library to integrate in other applications
* If it becomes relevant, compatibility priority order is Linux, BSD, Windows, Mac
* Weak copyleft. Tar-Sync itself belongs to the community; but can be integrated in proprietary applications.

## License

Lesser GPL, 3.0

This means you can link-against/embed it in a proprietary system without affecting the license of your own project.

If you do distribute a modified version of tar-sync itself however, you must release the modified source of your tar-sync copy to whomever should ask. This is limited to tar-sync code, and does not apply to any code that calls it or software that embeds it.
