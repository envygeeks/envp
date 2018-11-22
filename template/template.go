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
	"github.com/envygeeks/envp/logger"
)

// Template provides a context wrapper
// for all of our internal functions, for
// the template, and stuff for you.
type Template struct {
	use      string
	template *template.Template
	debug    bool
}

// New creates a new template, and logs it for
// the entire world to know if they really need to
// know what's going on for debugging purposes.
func New(debug bool) *Template {
	t := template.New("envp")
	helpers.New(t)

	return &Template{
		debug:    debug,
		template: t,
	}
}

// Use tells us to use this specific template
func (t *Template) Use(f NamedReader) {
	t.use = filepath.Base(f.Name())
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
	logger.Printf("attempting to parse %+v", r.Name())
	tt := t.template.New(filepath.Base(r.Name()))
	if b, err := ioutil.ReadAll(r); err != nil {
		logger.Fatalln(err)
	} else {
		if _, err := tt.Parse(string(b)); err != nil {
			logger.Fatalln(err)
		}
	}

	return tt
}

// Exec runs exec on the template.
// Before you hit this stage you should really be
// running Load(), and Parse() to get ready.
func (t *Template) Exec() []byte {
	var tt *template.Template

	if t.use != "" {
		logger.Printf("using requested %s", t.use)
		tt = t.template.Lookup(t.use)
		if tt == nil {
			logger.Fatalf("unable to find %s", t.use)
		}
	} else {
		templates := t.template.Templates()
		for _, ttt := range templates {
			if ttt.Name() == "base.gohtml" || ttt.Name() == "root.gohtml" {
				tt = ttt
				break
			}
		}

		if tt == nil {
			tt = templates[0]
			if tt == nil {
				logger.Fatalln("no template found")
			}
		}
	}

	b := &bytes.Buffer{}
	logger.Printf("executing %s", tt.Name())
	if err := tt.Execute(b, ""); err != nil {
		logger.Fatalln(err)
	}

	return b.Bytes()
}

// Write writes to stdout, or a file.
func (t *Template) Write(b []byte, w io.Writer) int {
	i, err := w.Write(b)
	if err != nil {
		logger.Fatalln(err)
	}

	return i
}
