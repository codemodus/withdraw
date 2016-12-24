package main

import (
	"fmt"
	"os"
)

func isFile(name string) bool {
	fi, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}

	return !fi.IsDir()
}

func main() {
	for _, name := range os.Args[1:] {
		if !isFile(name) {
			continue
		}

		if err := os.Remove(name); err != nil {
			fmt.Fprintf(os.Stdout,
				"error removing file %q: %s", name, err)
		}
	}
}
