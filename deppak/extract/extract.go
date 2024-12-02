package extract

import (
	"fmt"
	"io"
)

/*
 * Common functionality that all archivers might need
 */

func closeOrErr(ref io.Reader, message string) {
	if err := ref.Close(); err != nil {
		panic(fmt.Errorf(message) )
	}
}
