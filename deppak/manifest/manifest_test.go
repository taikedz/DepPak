package manifest

import (
    "testing"
    "fmt"
)

func Test_extractManifest(t *testing.T) {
    hash_val := "abcd1234"

    // NOTE - this is the old format. We need to expect a "deploy" key with a map of source to list of dest
    data, err := extractManifest(fmt.Sprintf(`{"url":"there", "hash":"%s", "dest":"mods/test_mod", "src":"subdir"}`, hash_val))

    if err != nil {
        t.Errorf("%s", err )
    }

    if data.Hash != hash_val {
        t.Errorf("Expected %s , got %s", hash_val, data.Hash)
    }
}
