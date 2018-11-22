// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"os"

	"github.com/envygeeks/envp/logger"
	"github.com/envygeeks/envp/template"
)

/**
 */
func main() {
	args := NewFlags().Parse()
	logger.SetOutput(ioutil.Discard)
	debug := args.Bool("debug")
	if debug {
		logger.SetOutput(os.Stderr)
	}

	ttemplate := template.New(debug)
	file, output, stdout := args.String("file"), args.String("output"), args.Bool("stdout")
	readers, writer := template.Open(file, output, stdout)
	defer template.Close(readers, writer)
	ttemplate.ParseFiles(readers)

	if len(readers) == 1 {
		ttemplate.Use(readers[0])
	}

	b := ttemplate.Exec()
	ttemplate.Write(b, writer)
}
