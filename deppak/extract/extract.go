package extract

import (
	"fmt"
	"io"
)

func UnpackArchive(archive_path str, dest_dir string, deploy_targets map[string][]string) error {
	// Detect archive type, and dispatch correctly
}


/*
 * Common functionality that all archivers might need
 */

func closeOrErr(ref io.Reader, message string) {
	if err := ref.Close(); err != nil {
		panic(fmt.Errorf(message) )
	}
}
