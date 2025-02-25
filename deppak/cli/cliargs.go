package cli

import (
    "flag"
    "fmt"
    "os"
	nstatus "net.taikedz.deppak/deppak/names"
    "net.taikedz.deppak/deppak/util"

    "github.com/taikedz/goargs/goargs"
)


type DepPakArgs struct {
	Manifest_path string
    Unpack_root string
}

func ParseCliArgs() DepPakArgs {
    var unpack_root string
    var manifest string

    parser := goargs.NewParser("DepPak options")
    parser.StringVar(&unpack_root, "unpack-to", "./", "Top level directory to unpack to")
    parser.Parse()
    if err := goargs.UnpackExactly(parser.Args(), &manifest); err != nil {
        util.Fail(nstatus.ERR_ARGUMENT_ERROR, "Expected one argument (MANIFEST)")
    }

    pos_info, err := os.Stat(manifest)
    if err != nil { util.Fail(nstatus.ERR_ARGUMENT_ERROR, "Could not access '%s': %s", manifest, err); }
    if pos_info.IsDir() { util.Fail(nstatus.ERR_ARGUMENT_ERROR, "'%s' is a directory, file required", manifest) }

    urt_info, err := os.Stat(unpack_root)
    if err != nil { util.Fail(nstatus.ERR_ARGUMENT_ERROR, "Could not access '%s': %s", unpack_root, err); }
    if !urt_info.IsDir() { util.Fail(nstatus.ERR_ARGUMENT_ERROR, "'%s' is not a directory", unpack_root) }

    return DepPakArgs{manifest, unpack_root}
}
