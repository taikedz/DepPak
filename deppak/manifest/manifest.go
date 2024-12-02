package manifest

import (
    "encoding/json"
    "fmt"
)

type SyncItem struct {
    Url  string `json:url`
    Hash string `json:hash`
    Dest string `json:dest`
    Src  string `json:src`
}

func LoadManifest(path string) (*SyncItem, error) {
    // NOTE - when loading the manifest, ensure there are no duplicate hashes (excl "-")
    // TODO - read from file
    fmt.Println(path)
    return extractManifest(`{"url":"there", "hash":"abcd1234", "dest":"mods/test_mod", "src":"subdir"}`)
}


func extractManifest(json_data string) (*SyncItem, error) {
    var data *SyncItem

    err := json.Unmarshal([]byte(json_data), &data)

    if err != nil {
        return nil, err
    }

    return data, nil
}
