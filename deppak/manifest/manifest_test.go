package manifest

import (
    "testing"
    "fmt"
)

func Test_extractManifest(t *testing.T) {
    hash_val := "abcd1234"
    dest_val := "mods/test_mod"

    // NOTE - this is the old format. We need to expect a "deploy" key with a map of source to list of dest
    datalist, err := extractManifest(fmt.Sprintf(`[{"url":"there", "hash":"%s", "dest":"%s", "src":"subdir"}]`, hash_val, dest_val))

    if err != nil {
        t.Errorf("%s", err )
    }

    data := datalist[0]
    if data.Hash != hash_val {
        t.Errorf("Full: Expected %s , got %s", hash_val, data.Hash)
    }

    if data.Dest != dest_val {
        t.Errorf("Full: Expected %s, got '%s'", dest_val, data.Dest)
    }

    // Omitting dest as it is optional
    datalist, err = extractManifest(fmt.Sprintf(`[{"url":"there", "hash":"%s", "src":"subdir"}]`, hash_val))

    if err != nil {
        t.Errorf("%s", err )
    }

    data = datalist[0]
    if data.Dest != "." {
        t.Errorf("Part: Expected Dest='%s', got '%s'", ".", data)
    }

    // Extract with missing fields
    datalist, err = extractManifest(`[{"url":"oops"}]`)
    if err == nil {
        t.Errorf("Missing: Should have failed on incomplete input! -> %s", datalist)
    }

    /*
    // This actually succeeds - unaccounted-for keys are simply ignored.
    datalist, err = extractManifest(`{"url":"ok", "hash":"what", "src": "here", "damn":"oops"}`)
    if err == nil {
        t.Errorf("Extra: Should have failed on bad input! -> %s", datalist)
    } else {
        fmt.Printf("Extra: Got expected Err = %s\n", err)
    }
    // */
    
    // Nulls
    datalist, err = extractManifest(`{"url":null}`)
    if err == nil {
        t.Errorf("Null: Should have failed on null data input! -> %s", datalist)
    }
}
