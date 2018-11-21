// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"io/ioutil"

	"github.com/envygeeks/envp/template"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Disable log output when we are testing.
	if v := flag.Lookup("test.v"); v != nil || v.Value.String() == "true" {
		log.SetOutput(ioutil.Discard)
	}
}

/**
 */
func main() {
	args := NewFlags().Parse()
	debug := args.Bool("debug")
	lvl := log.WarnLevel
	if debug {
		lvl = log.DebugLevel
	}

	log.SetLevel(lvl)
	ttemplate := template.New(debug)
	file, output, stdout := args.String("file"), args.String("output"), args.Bool("stdout")
	readers, writer := template.Open(file, output, stdout)
	defer template.Close(readers, writer)
	ttemplate.ParseFiles(readers)

	if len(readers) == 1 {
		ttemplate.Use(readers[0].Name())
	}

	b := ttemplate.Exec()
	ttemplate.Write(b,
		writer)
}
