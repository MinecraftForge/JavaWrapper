package util

import (
    "io"
    "github.com/lxq/lzma"
)

func DecompFile(file string)  {

}

func decompress(in io.Reader) io.Reader  {
   return lzma.NewReader(in)
}