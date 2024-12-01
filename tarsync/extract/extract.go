import (
	"fmt"
	"io"
)

func closeOrErr(ref io.Reader, message string) {
	if err := ref.Close(); err != nil {
		panic(fmt.Errorf(message) )
	}
}