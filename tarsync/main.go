package main

import (
    "fmt"
    "os"
    "sync"
)

const TARBALL_STORE = "~/.local/var/tarsync/tarballs"

func main() {
    args := argue.ParseCliArgs()
    printIfHelpFlag(args)

    all_entries := manifest.LoadManifest(args.jsonfile)

    var wg sync.WaitGroup
    wg.Add(len(all_entries))
    failures = make(chan string, len(all_entries))

    ensure_dir(TARBALL_STORE)

    for entry in all_entries {
        go download_entry(entry, &wg, failures)
    }

    wg.Wait()
    failures.Close() // FIXME - does this remove access to contents?

    var failure_strings string[]
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

    for entry := range all_entries {
        // Do not do this as concurrent - process in file declaration order
        extract_entry(entry)
    }
}

func download_entry(entry manifest.TarballEntry, wg *sync.WaitGroup, failures chan string) {
    defer wg.Done()

    // STEPS
    // - if tarball at hash does not exist
    //     - download to folder using hash string as name
    // - produce hash of tarball
    // - validate hash
    //     - if expected hash is "-" then print the URL and the computed hash
    // - if invalid (including "-"), write URL of failed item to failuures channel
}

func extract_entry(entry manifest.TarballEntry, destination_root string) {
    // Entry has: hash, url, dest, optional src

    // STEPS
    // - unpack tarball or tarball src/ target, into destination
}

