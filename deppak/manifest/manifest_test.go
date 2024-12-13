package manifest

import (
    "testing"
    "fmt"
)

func Test_extractManifest(t *testing.T) {
    hash_val := "abcd1234"
    dest_val := "mods/test_mod"

    // NOTE - this is the old format. We need to expect a "deploy" key with a map of source to list of dest
    data, err := extractManifest(fmt.Sprintf(`{"url":"there", "hash":"%s", "dest":"%s", "src":"subdir"}`, hash_val, dest_val))

    if err != nil {
        t.Errorf("%s", err )
    }

    if data.Hash != hash_val {
        t.Errorf("Expected %s , got %s", hash_val, data.Hash)
    }

    if data.Dest != dest_val {
        t.Errorf("Expected %s, got '%s'", dest_val, data.Dest)
    }

    // Omitting dest as it is optional
    data2, err2 := extractManifest(fmt.Sprintf(`{"url":"there", "hash":"%s", "src":"subdir"}`, hash_val))

    if err2 != nil {
        t.Errorf("%s", err2 )
    }

    if data2.Dest != "." {
        t.Errorf("Expected %s, got '%s'", dest_val, data2.Dest)
    }

    data3, err3 := extractManifest(`{"url":"oops"}`)
    if err3 == nil {
        t.Errorf("Should have failed on incomplete input! -> %s", data3)
    }

    /*
    // This actually succeeds - unaccounted-for keys are simply ignored.
    data4, err4 := extractManifest(`{"url":"ok", "hash":"what", "src": "here", "damn":"oops"}`)
    if err4 == nil {
        t.Errorf("Should have failed on bad input! -> %s", data4)
    } else {
        fmt.Printf("Got expected Err = %s\n", err4)
    }
    // */
    
    data5, err5 := extractManifest(`{"url":null}`)
    if err5 == nil {
        t.Errorf("Should have failed on null data input! -> %s", data5)
    } else {
        fmt.Printf("Got expected Err = %s\n", err5)
    }
}
