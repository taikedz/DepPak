package manifest

import (
    "encoding/json"
    "errors"
    "io/ioutil"
)

type SyncItem struct {
    Url  string
    Hash string
    Dest string
    Src  string
}

func LoadManifest(path string) ([]SyncItem, error) {
    // NOTE - when loading the manifest, ensure there are no duplicate hashes (excl "-")
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    // now we need to extract... a LIST of SyncItems ...
    // We might need to load the file to some interface{} , assert it is some sort of list
    //   and then coerce each into a SyncItem ...
    return extractManifest(string(content))
}


func extractManifest(json_data string) ([]SyncItem, error) {
    var data []SyncItem

    // Interestingly:
    //   unknown keys are ignored (value discarded)
    //   missing keys are non-populated
    //   null-values are non-populated when target is unmarshalled as SyncItem,
    //      but actually fails when []SyncItem...
    // Non-populated values leave the struct at default values
    // So null in JSON leads to a struct string field becoming  ""
    err := json.Unmarshal([]byte(json_data), &data)

    if err != nil {
        return nil, err
    }

    for i, item := range data {
        if item.Dest == "" {
            item.Dest = "."
        }

        if item.Url == "" {return nil, errors.New("Missing url")}
        if item.Src == "" {return nil, errors.New("Missing src")}
        if item.Hash == "" {return nil, errors.New("Missing hash")}

        // range over []struct produces copies.
        // re-assign back
        // https://stackoverflow.com/a/51106020/2703818
        data[i] = item
    }

    return data, nil
}
