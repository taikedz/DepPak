package cli

import (
    "os"
    "fmt"
)


const HELP_STR =`
deppak [--unpack-root=DIRPATH] MANIFEST

Call deppak and supply a manifest file.

If DIRPATH is specified, unpack files to that location.
`

func PrintIfHelpFlag() {
    /* Print the overall help string and exit, if "--help" found in CLI tokens
     */
    for _, token := range os.Args {
        if token == "--help" {
            fmt.Printf("%s\n", HELP_STR)
            os.Exit(0)
        }
    }
}
