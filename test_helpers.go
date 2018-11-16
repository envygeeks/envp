// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func captureOut(f func()) string {
	var s strings.Builder
	log.SetOutput(&s)
	f()

	log.SetOutput(os.Stderr)
	return s.String()
}

func createTemplate(data string) *Template {
	a := NewFlags().Parse()
	t1, _ := ioutil.TempFile("", "")
	t2, _ := ioutil.TempFile("", "")
	defer t1.Close()
	defer t2.Close()

	a["file"], a["output"] = t1.Name(), t2.Name()
	t := NewTemplate(&a)
	if data != "" {
		t.template.Parse(data)
	}

	return t
}
