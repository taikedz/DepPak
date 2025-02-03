package main

import (
    "fmt"
    "os"
    "sync"

    "net.taikedz.deppak/deppak/util"
    "net.taikedz.deppak/deppak/manifest"
    "net.taikedz.deppak/deppak/cli"
    "net.taikedz.deppak/deppak/names"
    "net.taikedz.deppak/deppak/net"
    "net.taikedz.deppak/deppak/cache"
)


func main() {
    cli.PrintIfHelpFlag()
    args := cli.ParseCliArgs()

    all_entries, err := manifest.LoadManifest(args.Manifest_path)
    if err != nil {
        fmt.Println(err)
        os.Exit(names.ERR_BAD_MANIFEST)
    }

    var wg sync.WaitGroup
    wg.Add(len(all_entries))
    failures := make(chan string, len(all_entries))

    for _, entry := range all_entries {
        go func() {
            defer wg.Done()
            download_entry(entry, failures)
        }()
    }

    wg.Wait()
    close(failures)

    var failed = false

    if failed {
        // iterate fails - if failures exist, print all  failures
        // then exit without unpacking
        for fail_entry := range failures {
            fmt.Println(fail_entry)
            failed = true
        }

        os.Exit(names.ERR_FAILURES)
    }
    for _, entry := range all_entries {
        // DO NOT do this as concurrent - process in file declaration order
        // Allows progressive overwriting
        fmt.Printf("Extracting %v to %s...\n", entry, destination_root)
        extract.ExtractZip(cache.GetFileFor(entry.Hash), args.Unpack_root, entry.Deploy)
    }
}

func download_entry(entry manifest.Dependency, failures chan string) {
    fmt.Printf("Downloading %v ...\n", entry)
    if ! cache.Exists(entry.Hash) {
        temp_filepath, err := net.FetchHttp(entry.Url)
        hash, _ := cache.Retain(temp_filepath)
        if hash != entry.Hash {
            fmt.Printf("File from %s has hash %s ; expected %s\n", entry.Url, hash, entry.Hash)
            failures <- filepath
        }
    } else {
        filepath := cache.GetFileFor(entry.Hash)
        hash := util.HashFile(filepath)
        if hash != entry.Hash {
            fmt.Printf("/!\\ Cached file for hash '%s' has actual hash %s\n\t(cached file: %s).\n\t---> Tampering detected?\n", entry.hash, hash, filepath)
            fmt.Printf("Purge=%s to resolve.\n", hash) // make this line easily parsable for CI systems: `^Purge=(\S+)`
            failures <- filepath
        }
    }
}
