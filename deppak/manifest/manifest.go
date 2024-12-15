package manifest

import (
    "encoding/json"
    "errors"
    "io/ioutil"
)

type Dependency struct {
    Hash string
    Url string
    Deploy map[string][]string
}


func LoadManifest(path string) ([]Dependency, error) {
    // NOTE - when loading the manifest, ensure there are no duplicate hashes (excl "-")
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    return extractManifest(string(content))
}


func extractManifest(json_data string) ([]Dependency, error) {
    var data []Dependency

    // Interestingly:
    //   unknown keys are ignored (value discarded)
    //   missing keys are non-populated
    //   null-values are non-populated when target is unmarshalled as SyncItem,
    //      but actually fails when []SyncItem...
    // Non-populated values leave the struct at default values
    // So null in JSON leads to a struct string field becoming  ""
    err := json.Unmarshal([]byte(json_data), &data)

    if err != nil {
        return nil, errors.New("JSON parsing error: "+err.Error())
    }

    for _, item := range data {
        if item.Url == "" {return nil, errors.New("Missing url")}
        if item.Hash == "" {return nil, errors.New("Missing hash")}
    }

    return data, nil
}
