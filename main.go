package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	cFS := compoundFS{"base", "overrides"}

	fs := http.FileServer(cFS)
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3002", nil)
}

type compoundFS struct {
	base, overrides string
}

func (c compoundFS) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}

	attempts := []string{c.overrides, c.base}

	for i, attempt := range attempts {
		dir := string(attempt)
		if dir == "" {
			dir = "."
		}
		f, err := os.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
		if err != nil {
			log.Println("error on base: ", err)
			if i == len(attempts) {
				return nil, err

			}
			continue
		}
		return f, nil
	}
	return nil, errors.New("Something went wrong")
}
