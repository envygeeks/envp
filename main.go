// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"os"

	"github.com/envygeeks/envp/flags"
	"github.com/envygeeks/envp/logger"
	upstream "github.com/envygeeks/envp/template"
)

/**
 */
func main() {
	args := flags.Parse()
	logger.SetOutput(ioutil.Discard)
	debug := args.Bool("debug")
	if debug {
		logger.SetOutput(os.Stderr)
	}

	template := upstream.New(debug)
	file, output := args.String("file"), args.String("output")
	readers, writer := upstream.Open(file, output)
	defer upstream.Close(readers, writer)
	template.ParseFiles(readers)

	if len(readers) == 1 {
		template.Use(readers[0])
	}

	b := template.Compile()
	template.Write(b, writer)
}
