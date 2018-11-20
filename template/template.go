// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/envygeeks/envp/helpers"
	log "github.com/sirupsen/logrus"
)

// Template provides a context wrapper
// for all of our internal functions, for
// the template, and stuff for you.
type Template struct {
	template *template.Template
	debug    bool
}

// New creates a new template, and logs it for
// the entire world to know if they really need to
// know what's going on for debugging purposes.
func New(debug bool) *Template {
	externalTemplate := template.New("envp")
	helpers.New(externalTemplate)
	return &Template{
		template: externalTemplate,
		debug:    debug,
	}
}

// ParseFiles parses all your readers
func (t *Template) ParseFiles(fs []*os.File) []*template.Template {
	var ts []*template.Template
	for _, v := range fs {
		ts = append(ts, t.Parse(v))
	}

	return ts
}

// Parse parses the templates.
func (t *Template) Parse(r NamedReader) *template.Template {
	log.Printf("attempting to parse %+v", r.Name())
	tt := t.template.New(filepath.Base(r.Name()))
	if b, err := ioutil.ReadAll(r); err != nil {
		log.Fatalln(err)
	} else {
		if _, err := tt.Parse(string(b)); err != nil {
			log.Fatalln(err)
		}
	}

	return tt
}

// Exec runs exec on the template.
// Before you hit this stage you should really be
// running Load(), and Parse() to get ready.
func (t *Template) Exec() []byte {
	var tt *template.Template

	templates := t.template.Templates()
	for _, ttt := range templates {
		if ttt.Name() == "base.gohtml" || ttt.Name() == "root.gohtml" {
			tt = ttt
			break
		}
	}

	tt = templates[0]
	if tt == nil {
		log.Fatalln("unable to find a template")
	}

	var b bytes.Buffer
	log.Infof("executing %s", tt.Name())
	if err := tt.Execute(&b, ""); err != nil {
		log.Fatal(err)
	}

	return b.Bytes()
}

// Write writes to stdout, or a file.
func (t *Template) Write(b []byte, w io.Writer) int {
	i, err := w.Write(b)
	if err != nil {
		log.Fatalln(err)
	}

	return i
}
