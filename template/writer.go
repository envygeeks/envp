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

// Writer interface
type Writer interface {
	Close() error
	io.Writer
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
