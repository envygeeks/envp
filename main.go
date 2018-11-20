// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"github.com/envygeeks/envp/template"
	log "github.com/sirupsen/logrus"
)

/**
 */
func main() {
	a := NewFlags().Parse()
	d := a.Bool("debug")
	l := log.WarnLevel
	if d {
		l = log.DebugLevel
	}

	log.SetLevel(l)
	f := a.String("file")
	o := a.String("output")
	s := a.Bool("stdout")

	r, w := template.Open(f, o, s)
	defer template.Close(r, w)
	t := template.New(d)
	t.ParseFiles(r)

	if len(r) == 1 {
		t.Use(r[0].Name())
	}

	b := t.Exec()
	t.Write(b, w)
}
