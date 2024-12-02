package main

func test_extractManifest() {
    data, err := extractManifest(`{"url":"there", "hash":"abcd1234", "dest":"mods/test_mod", "src":"subdir"}`)

    if err != nil {
        log.Fatal(err)
    }

    // do asserttions here
    fmt.Println(data)
    fmt.Println(data.Hash)
}
