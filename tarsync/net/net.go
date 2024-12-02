package net

import (
    "os"
    "net/http"
    "io"
)

// Apparently the following code does stream data,
// rather than load to memory, then dump to file....
// https://www.golangcode.com/download-a-file-from-a-url/

func FetchHttp(url string, destfile string) (err error) {
  out, err := os.Create(destfile)
  if err != nil  {
    return err
  }
  defer out.Close()

  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

// TODO - also support SSH/SCP
// Ignore: FTP(S), Telnet, etc. These are defunct protocols that no longer have reason to be used.
// also support `file://` scheme which should create a symlink
