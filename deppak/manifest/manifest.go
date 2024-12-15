package manifest

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
)

type Dependency struct {
    Hash string
    Url string
    Deploy map[string][]string
}


func LoadManifest(path string) ([]Dependency, error) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    dependencies, xerr := extractManifest(string(content))
    if xerr != nil {
        return nil, xerr
    }

    if dupes := findHashDuplicates(dependencies); len(dupes) > 0 {
        return nil, errors.New(fmt.Sprintf("Duplicate hashes found - assemble all under one hash: %s", dupes))
    }

    // TODO - set default for "deploy", and check for illegal path skirting

    return dependencies, xerr
}

func findHashDuplicates(dependencies []Dependency) []string {
    var seen_hashes = make(map[string]struct{}, 0)
    var dupes_map = make(map[string]struct{}, 0)
    var SEEN struct{}

    for _, dep := range dependencies {
        if dep.Hash == "-" { continue; }

        if _, found := seen_hashes[dep.Hash]; found {
            dupes_map[dep.Hash] = SEEN
        } else {
            seen_hashes[dep.Hash] = SEEN
        }
    }

    var dupes_list = make([]string, 0, len(dupes_map))
    for hash,_ := range(dupes_map) {
        dupes_list = append(dupes_list, hash)
    }

    return dupes_list
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
        return nil, err
    }

    for _, item := range data {
        if item.Url == "" {return nil, errors.New("Missing url")}
        if item.Hash == "" {return nil, errors.New("Missing hash")}
    }

    return data, nil
}
