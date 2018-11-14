// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func parseTemplate(file, output string, glob bool, stdout bool) {
	var tt *template.Template
	t := template.New("EnvPTemplate")
	t.Funcs(funcs(t))

	var fs []string
	if glob {
		var err error
		p := filepath.Join(file, "*.gohtml")
		fs, err = filepath.Glob(p)
		if err != nil {
			log.Fatalln(err)
		}
	} else {

		// This might be dumb, we could always just
		// leave it as a string, and then detect that
		// futher down the line?!
		fs = []string{
			file,
		}
	}

	// Parse them and ship along the errors, it's not
	// really our concern if there is an error as it's
	// a downstream error that should be risen up...
	if _, err := t.ParseFiles(fs...); err != nil {
		log.Printf("Unable to parse the files %+v", fs)
		log.Fatalln(err)
	}

	/**
	 */
	if len(fs) == 1 {
		name := filepath.Base(fs[0])
		if tt = t.Lookup(name); tt == nil {
			log.Fatalf("no template \"%s\"", name)
		}
	} else {
		for _, v := range []string{"base.gohtml", "root.gohtml"} {
			tt = t.Lookup(v)
			if tt != nil {
				break
			}
		}

		if tt == nil {
			log.Fatalln("no base, or root")
		}
	}

	var s strings.Builder
	tt.Execute(&s, "")
	if !stdout {
		d := filepath.Dir(output)
		if err := os.MkdirAll(d, 0644); err != nil {
			log.Fatalln(err)
		} else {
			b := []byte(s.String())
			if err := ioutil.WriteFile(output, b, 0644); err != nil {
				log.Fatalln(err)
			}
		}
	} else {
		fmt.Println(s.String())
	}
}

/**
 */
func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

/**
 */
func main() {
	a := flags()
	g := a.Bool("glob")
	o := expand(a.String("output"))
	f := expand(a.String("file"))
	s := a.Bool("stdout")

	if isExist(f) {
		parseTemplate(f, o,
			g, s)
	} else {
		// Maybe I should make this more simple?
		log.Fatalf("-file=\"%s\" doesn't exist", f)
		os.Exit(1)
	}
}
