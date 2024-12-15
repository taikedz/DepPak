package manifest

import (
    "testing"
    "fmt"
    // "os"
    // "encoding/json"
    "reflect"
)

func Test_extractManifest(t *testing.T) {
    hash_val := "abcd1234"
    dest_val := "mods/test_mod"

    datalist, err := extractManifest(fmt.Sprintf(`[{"url"/:"there", "hash":"%s", "deploy":{"src":["%s"]} }]`, hash_val, dest_val))

    if err != nil {
        t.Errorf("%s : %s", reflect.TypeOf(err), err)
        // if _, type_found := err.(json.InvalidUnmarshalError); type_found {
        //     os.Exit(1)
        // }
    }

    data := datalist[0]
    if data.Hash != hash_val {
        t.Errorf("Full: Expected %s , got %s", hash_val, data.Hash)
    }

    if data.Deploy["src"][0] != dest_val {
        t.Errorf("Full: Expected %s, got '%s'", dest_val, data.Deploy)
    }

    // Omitting `deploy` as it is optional
    datalist, err = extractManifest(fmt.Sprintf(`[{"url":"there", "hash":"%s"}]`, hash_val))

    if err != nil {
        t.Errorf("%s", err )
    }

    data = datalist[0]
    if len(data.Deploy) != 0 {
        t.Errorf("Part: Expected empty Deploy, got '%s'", data)
    }

    // Extract with missing fields
    datalist, err = extractManifest(`[{"url":"oops"}]`)
    if err == nil {
        t.Errorf("Missing: Should have failed on incomplete input! -> %s", datalist)
    }

    // This actually succeeds - unaccounted-for keys are simply ignored.
    datalist, err = extractManifest(`[{"url":"ok", "hash":"what", "src": "here", "damn":"oops"}]`)
    if err != nil {
        t.Errorf("Extra: Should have succeded ignoring extraneous input! -> %s", datalist)
    }
    
    // Nulls
    datalist, err = extractManifest(`{"url":null}`)
    if err == nil {
        t.Errorf("Null: Should have failed on null data input! -> %s", datalist)
    }
}
