package util

import (
	"fmt"
	"os"
)

func Fail(status int, message string, tokens... any) {
    fmt.Println(fmt.Sprintf(message, tokens...))
    os.Exit(status)
}


func HashOfFile(filepath string) {
}
