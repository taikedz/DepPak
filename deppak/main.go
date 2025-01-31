package main

import (
    "fmt"
    "os"
    "sync"

    "net.taikedz.deppak/deppak/manifest"
    "net.taikedz.deppak/deppak/cli"
    "net.taikedz.deppak/deppak/names"
//    "net.taikedz.deppak/deppak/net"
)

const ARCHIVE_STORE = "~/.local/var/deppak/z"

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

    os.MkdirAll(ARCHIVE_STORE, 0700)

    for _, entry := range all_entries {
        go func() {
            defer wg.Done()
            download_entry(entry, failures)
        }()
    }

    wg.Wait()
    close(failures)

    var failure_strings []string
    _ = failure_strings // DEBUG
    var failed = false

    if failed {
        // iterate fails - if failures exist, print all  failures
        // then exit without unpacking
        for fail_entry := range failures {
            fmt.Println(fail_entry)
            failed = true
        }

        os.Exit(1)
    }
    for _, entry := range all_entries {
        // Do not do this as concurrent - process in file declaration order
        extract_entry(entry, args.Unpack_root)
    }
}

func download_entry(entry manifest.Dependency, failures chan string) {
    fmt.Printf("Downloading %v ...\n", entry)
    _ = failures
    // STEPS
    // - if tarball at hash does not exist
    //     - download to folder using hash string as name
    // - produce hash of tarball
    // - validate hash
    //     - if expected hash is "-" then print the URL and the computed hash
    // - if invalid (including "-"), write URL of failed item to failuures channel
}

func extract_entry(entry manifest.Dependency, destination_root string) {
    fmt.Printf("Extracting %v to %s...\n", entry, destination_root)
    // Entry has: hash, url, dest, optional src

    // STEPS
    // - unpack tarball or tarball src/ target, into destination
}
