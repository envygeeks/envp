// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"io"
	"os"
	"path/filepath"

	"github.com/envygeeks/envp/logger"
)

// Reader interface
type Reader interface {
	Close() error
	Name() string
	io.Reader
}

func reader(file string) *os.File {
	logger.Printf("opening a reader to %s", file)
	reader, err := os.Open(file)
	if err != nil {
		logger.Fatalln(err)
	}

	return reader
}

func readers(file string) []Reader {
	file, err := filepath.Abs(file)
	if err != nil {
		logger.Fatalln(err)
	}

	finfo, err := os.Stat(file)
	if err == nil {
		if !finfo.IsDir() {
			reader := reader(file)
			return []Reader{
				reader,
			}
		}
	} else {
		logger.Fatalln(err)
	}

	files := []Reader{}
	logger.Printf("looking for *.gohtml in %s", file)
	p := filepath.Join(file, "*.gohtml")
	all, err := filepath.Glob(p)
	if err != nil {
		logger.Fatalln(err)
	} else {
		for _, v := range all {
			files = append(files, reader(v))
		}
	}

	return files
}
