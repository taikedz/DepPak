package manifest

import (
    "encoding/json"
    "fmt"
    "errors"
    "io/ioutil"
)

type SyncItem struct {
    Url  string `json:url`
    Hash string `json:hash`
    Dest string `json:dest`
    Src  string `json:src`
}

func LoadManifest(path string) (*SyncItem, error) {
    // NOTE - when loading the manifest, ensure there are no duplicate hashes (excl "-")
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    // now we need to extract... a LIST of SyncItems ...
    // We might need to load the file to some interface{} , assert it is some sort of list
    //   and then coerce each into a SyncItem ...
    return extractManifest(content)
}


func extractManifest(json_data string) (*SyncItem, error) {
    var data *SyncItem

    // Interestingly:
    //   unknown keys are ignored (value discarded)
    //   missing keys are non-populated
    //   null-values are non-populated
    // Non-populated values leave the struct at default values
    // So null in JSON leads to a struct string field becoming  ""
    err := json.Unmarshal([]byte(json_data), &data)

    if err != nil {
        return nil, err
    }

    if data.Dest == "" {
        data.Dest = "."
    }

    if data.Url == "" {return nil, errors.New("Missing url")}
    if data.Src == "" {return nil, errors.New("Missing src")}
    if data.Hash == "" {return nil, errors.New("Missing hash")}

    return data, nil
}
