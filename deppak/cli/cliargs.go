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

    if len(positionals) != 1 {
        fail(nstatus.ARGUMENT_ERROR, "Expecting one MANIFEST argument")
    }

    if token, ok := checkTrailingFlags(positionals); !ok {
        fail(nstatus.ARGUMENT_ERROR, "Found flag '%s' . Place all flags before positional arguments.\n", token)
    }

    pos_info, err := os.Stat(positionals[0])
    if err != nil { fail(nstatus.ARGUMENT_ERROR, "Could not access '%s': %s", positionals[0], err); }
    if pos_info.IsDir() { fail(nstatus.ARGUMENT_ERROR, "'%s' is a directory, file required", positionals[0]) }

    urt_info, err := os.Stat(unpack_root)
    if err != nil { fail(nstatus.ARGUMENT_ERROR, "Could not access '%s': %s", unpack_root, err); }
    if !urt_info.IsDir() { fail(nstatus.ARGUMENT_ERROR, "'%s' is not a directory", unpack_root) }

    return DepPakArgs{positionals[0], unpack_root}
}

func fail(status int, message string, tokens... any) {
    fmt.Println(fmt.Sprintf(message, tokens...))
    os.Exit(status)
}
