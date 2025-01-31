package cli

import (
    "flag"
    "fmt"
    "os"
	nstatus "net.taikedz.deppak/deppak/names"
)


type DepPakArgs struct {
	Manifest_path string
    Unpack_root string
}


func checkTrailingFlags(tokens []string) (token string, ok bool) {
    /* Check all tokens for anything that looks like a flag.
    If one is found, returns it with false
    Else, return empty string with true.

    Returns:
        token string - potentially faulty token
        ok bool - true if no flag found, false if unexpected flag found
    */
    for _, t := range tokens {
        if len(t) > 0 && t[0] == '-' {
            return t, false
        }
    }

    return "", true
}

func ParseCliArgs() DepPakArgs {
    /* Want to be able to either of
     *
     * deppak MANIFEST --unpack-root=./path
     * deppak --unpack-root=./path MANIFEST
     *
     * That is, flags can bloody well come after the positionals
     */
    var unpack_root string
    flag.StringVar(&unpack_root, "unpack-root", "./", "Top level directory to unpack to")
    flag.Parse()
    positionals := flag.Args()

    if len(positionals) == 0 {
        fmt.Printf("Expecting MANIFEST argument")
        os.Exit(nstatus.ARGUMENT_ERROR)
    }

    if token, ok := checkTrailingFlags(positionals); !ok {
        fmt.Printf("Found flag '%s' . Place all flags before positional arguments.\n", token)
        os.Exit(nstatus.ARGUMENT_ERROR)
    }

    return DepPakArgs{unpack_root, positionals[0]}
}
