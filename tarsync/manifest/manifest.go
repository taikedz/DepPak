package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type SyncItem struct {
    Url  string `json:url`
    Hash string `json:hash`
    Dest string `json:dest`
    Src  string `json:src`
}

func main() {
    ReadManifest("bogus")
}

func LoadManifest(path string) {
    // NOTE - when loading the manifest, ensure there are no duplicate hashes (excl "-")
    fmt.Println(path)

    data, err := extractManifest(`{"url":"there", "hash":"abcd1234", "dest":"mods/test_mod", "src":"subdir"}`)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(data)
    fmt.Println(data.Hash)
}


func extractManifest(json_data string) (*SyncItem, error) {
    var data *SyncItem

    err := json.Unmarshal([]byte(json_data), &data)

    if err != nil {
        return nil, err
    }

    return data, nil
}
