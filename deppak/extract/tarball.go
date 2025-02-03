package extract

// TODO : support other compressions: lzma and bzip2, and also no-compression

// from https://stackoverflow.com/a/57640231

import (
	"archive/tar"
	"io"
	"fmt"
)

func ExtractTarGz(archive_path str, dest_dir string, deploy_targets map[string][]string) err {
	stream, err := os.Open(aarchive_path)
	if err != nil {
		return err
	}
	defer closeOrErr(stream, "Could not extract file "+path)

	return extractTarGzStream(stream) // FIXME - pass down filters
}

func extractTarGzStream(gzip_stream io.Reader) error {
	uncompressed_stream, err := gzip.NewReader(gzip_stream)
	if err != nil {
		return err
	}

	tar_reader := tar.NewReader(uncompressed_stream)

	for true {
		header, err := tar_reader.Next()
		if err == io.EOF {
			break
		}
		else if err != nil {
			return err
		}

		switch header.Typeflag {
			case tar.TypeDir:
				if err := os.MkdirAll(header.Name, 0700) { // can we get the type of the inner file?
					return err
				}
			case tar.TypeReg:
				out_file, err := os.Create(header.Name)
				if err != nil {
					return err
				}
				defer closeOrErr(out_file, "Could not close "+header.Name)
				if _, err := io.Copy(out_file, tar_reader); err != nil {
					return err
				}
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("Could not extact %s : %s", header.TypeFlag, header.Name)
		}
	}

	return nil
}
