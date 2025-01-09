type DepPakArgs struct {
	manifest_path string
    unpack_root string
}

func ParseCliArgs() (DepPakArgs, error) {
    /* Want to be able to either of
     *
     * deppak MANIFEST --unpack-root=./path
     * deppak --unpack-root=./path MANIFEST
     *
     * That is, flags can bloody well come after the positionals
     */
}
