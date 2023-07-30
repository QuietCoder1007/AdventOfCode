package fileutils

import (
	"io"
	"log"
	"os"
)

func ReadTextFile(path string) io.Reader {
	// Open file
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	return f
}
