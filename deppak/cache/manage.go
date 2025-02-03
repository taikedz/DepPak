package cache


const ARCHIVE_STORE = "~/.local/var/deppak/z"


/* Move file to its hash location in cache dir
*/
func Retain(filepath strng) (hash string, err Error) {
    os.MkdirAll(ARCHIVE_STORE, 0700)
}

/* Determine existinece of tarball at hash location
Error if cannot read location
*/
func Exists(hash string) (exists bool, err Error) {
}

/* Get the filepath for the corresponding hash
Error if hash is not found, or file is not found where it was expected
*/
func GetFileFor(hash string) (filepath string, err Error) {
}