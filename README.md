# TarSync

A small utility inspired by the "dependencies" section in `build.zig.zon`.

Allows setting dependencies from a file pointing to tarballs, including hash-checking.

Reads a `tarsync.txt` file which has the following format:

```sh
# A top-level path where the following names will be stored
top      ./some/rel/path

# A path         MD5 hash      A URL to a tarball
name_in_top      abcd1234      https://some_url/file.tgz
# Will unpack the tarball directly to ./some/rel/path/name_in_top/

# A last option specifies the subfolder we are specifically interested in
#    storing at the named location
name_in_top      abcd1234      https://some_url/file.tgz   src/
# Will unpack the tarball, and move the contents of src/ into ./some/rel/path/name_in_top/
```

This syncs all the specified tarballs, validates the hash, and unpacks it to a location.

Hash can be specified as `-` to cause the downloaded item's hash to be printed to console. Unpack will not proceed in this case. Re-run with the hash populated to the file to proceed.

## Applications

This can be used where any bunch of distributable files is to be collected.

I would have happily used this for distributing a spec list of Minetest mods for a server, for example.

Any language, project or application that doesn't have its own bundle distribution system can use this.

As a libaray, this can be integrated into an application for it to provide extensibility with user plugins.

## Goals

* Simple specification of file
* Standalone commandline utility as single binary with no runtime dependencies
* Available as linkable library for Zig and C
* Compatibility priority order Linux, BSD, Windows, Mac
* Weak copyleft. Tar-Sync itself belongs to the community; but can be integrated in proprietary applications.

## License

Lesser GPL, 3.0

This means you can link against it in a proprietary system without affecting the license of your own project.

If you do distribute a modified version of tar-sync itself however, you must release the modified source of your tar-sync copy to whomever should ask. This is limited to tar-sync code, and does not apply to any code that calls it.
