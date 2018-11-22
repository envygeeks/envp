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

// NamedReader is a reader with Name()
// that way it can be a full file, or otherwise
// just as long as it has Name(), and Read()
type NamedReader interface {
	io.Reader
	Name() string
}

func reader(file string) *os.File {
	logger.Printf("opening a reader to %s", file)
	reader, err := os.Open(file)
	if err != nil {
		logger.Fatalln(err)
	}

	return reader
}

func readers(file string) []*os.File {
	file, err := filepath.Abs(file)
	if err != nil {
		logger.Fatalln(err)
	}

	finfo, err := os.Stat(file)
	if err == nil {
		if !finfo.IsDir() {
			reader := reader(file)
			return []*os.File{
				reader,
			}
		}
	} else {
		logger.Fatalln(err)
	}

	files := []*os.File{}
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

func writer(file string, stdout bool) *os.File {
	var fm os.FileMode
	if stdout {
		logger.Println("using stdout")
		return os.Stdout
	}

	file, err := filepath.Abs(file)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Printf("opening a writer to %s", file)
	fm, op := 0644, os.O_CREATE|os.O_WRONLY
	writer, err := os.OpenFile(file, op, fm)
	if err != nil {
		logger.Fatalln(err)
	}

	return writer
}
