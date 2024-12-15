package manifest

import (
    "testing"
    "fmt"
    "errors"
    "encoding/json"
)


// Parse error checking should not be needed when the JSON is hardcoded
// But I did hard code bad JSON at one point and tore my hair out
// I also tore lots of hair out to find out how to detect error types
// It's insane, so using a type declaration workaround
//   and documenting this for myself...

// Create a throwaway variable of the type of the error to detect
// then use `errors.As(err, &ParseErr)`
var ParseErr *json.SyntaxError


func Test_extractManifest(t *testing.T) {
    hash_val := "abcd1234"
    dest_val := "mods/test_mod"
    
    datalist, err := extractManifest(fmt.Sprintf(`[{"url":"there", "hash":"%s", "deploy":{"src":["%s"]} }]`, hash_val, dest_val))
    if err != nil {
        t.Errorf("%s", err)
    }
    if errors.As(err, &ParseErr) { return ; }

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
    if errors.As(err, &ParseErr) { return ; }

    data = datalist[0]
    if len(data.Deploy) != 0 {
        t.Errorf("Part: Expected empty Deploy, got '%s'", data)
    }

    // Extract with missing fields
    datalist, err = extractManifest(`[{"url":"oops"}]`)
    if err == nil {
        t.Errorf("Missing: Should have failed on incomplete input! -> %s", datalist)
    }
    if errors.As(err, &ParseErr) { return ; }

    // This actually succeeds - unaccounted-for keys are simply ignored.
    datalist, err = extractManifest(`[{"url":"ok", "hash":"what", "src": "here", "damn":"oops"}]`)
    if err != nil {
        t.Errorf("Extra: Should have succeded ignoring extraneous input! -> %s", datalist)
    }
    if errors.As(err, &ParseErr) { return ; }
    
    // Nulls
    datalist, err = extractManifest(`{"url":null}`)
    if err == nil {
        t.Errorf("Null: Should have failed on null data input! -> %s", datalist)
    }
    if errors.As(err, &ParseErr) { return ; }
}


func Test_findHashDuplicates(t *testing.T) {
    var deps = make([]Dependency, 0, 5)

    deps = append(deps, Dependency{"abc", "here", nil})
    deps = append(deps, Dependency{"def", "there", nil})
    deps = append(deps, Dependency{"-", "this", nil})
    deps = append(deps, Dependency{"-", "that", nil})

    if dupes := findHashDuplicates(deps); len(dupes) > 0 {
        t.Errorf("Should not have found duplicates, found %s", dupes)
    }

    deps = append(deps, Dependency{"abc", "here", nil})

    dupes := findHashDuplicates(deps)
    if len(dupes) != 1 {
        t.Errorf("Should have found one duplicate, found %s", dupes)
    } else {
        if dupes[0] != "abc" {
            t.Errorf("Expected 'abc', found %s", dupes)
        }
    }

}
