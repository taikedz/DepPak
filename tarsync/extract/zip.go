// less stumbly adaptation of https://stackoverflow.com/a/24792688

package extract

import (
	"archive/zip"
	"os"
	"path/filepath"
	"strings"
)

func ExtractZip(archive_path str, dest_dir string) error { // FIXME - add a path filter
	archive_handler, err := zip.OpenReader(archive_path)
	if err != nil {
		return err
	}
	defer closeOrErr(archive_handler, "Filed to close archive_path")

	os.MkdirAll(dest_dir, 0700)

	for _, inner_file := range archive_handler.File {
		err := extractInnerFile(inner_file, dest_dir) {
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func extractInnerFile(inner_file *zip.File, dest_dir string) error {
	src_handler, err := inner_file.Open()
	if err != nil {
		return err
	}
	defer closeOrErr(src_handler, "Failed to close reader for <inner file>")

	path := filepath.Join(dest_dir, inner_file.Name)

	if !strings.HasPrefix(path, filepath.Clean(dest_dir)+string(os.PathSeparator)) {
		return fmt.Errorf("Illegal file path: %s", path)
	}

	if inner_file.FileInfo().IsDir() {
		os.MkdirAll(path, inner_file.Mode())
	} else {
		os.MkdirAll(filepath.Dir(path), inner_file.Mode())

		outer_file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, inner_file.Mode())

		if err != nil {
			return err
		}
		defer closeOrErr(outer_file, "Failed to finalize <dest file>")

		_,err = io.Copy(outer_file, src_handler)
		if err != nil {
			return err
		}
	}
	retrn nil
}
