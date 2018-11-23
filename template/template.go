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

	"github.com/envygeeks/envp/logger"
	"github.com/envygeeks/envp/template/helpers"
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
func (t *Template) Use(f Reader) {
	t.use = filepath.Base(f.Name())
}

// ParseFiles parses all your readers
func (t *Template) ParseFiles(fs []Reader) []*template.Template {
	var ts []*template.Template
	for _, v := range fs {
		if r, ok := v.(Reader); ok {
			ts = append(ts, t.Parse(r))
		}
	}

	return ts
}

// Parse parses the templates.
func (t *Template) Parse(r Reader) *template.Template {
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

// Writer interface
type Writer interface {
	Name() string
	Close() error
	io.Writer
}

// Write writes to stdout, or a file.
func (t *Template) Write(b []byte, w Writer) int {
	i, err := w.Write(b)
	if err != nil {
		logger.Fatalln(err)
	}

	return i
}

func writer(file string) *os.File {
	var fm os.FileMode
	if file == "" {
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

// Reader interface
type Reader interface {
	Close() error
	Name() string
	io.Reader
}

func reader(file string) *os.File {
	logger.Printf("opening a reader to %s", file)
	reader, err := os.Open(file)
	if err != nil {
		logger.Fatalln(err)
	}

	return reader
}

func readers(file string) []Reader {
	file, err := filepath.Abs(file)
	if err != nil {
		logger.Fatalln(err)
	}

	finfo, err := os.Stat(file)
	if err == nil {
		if !finfo.IsDir() {
			reader := reader(file)
			return []Reader{
				reader,
			}
		}
	} else {
		logger.Fatalln(err)
	}

	files := []Reader{}
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

// Open opens all the readers, and writers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Open(r, w string) (_readers []Reader, _writer Writer) {
	return readers(r), writer(w)
}

/**
 */
func cWriter(w Writer) {
	if w != os.Stdout {
		w.Close()
	}
}

/**
 */
func cReader(r []Reader) {
	for _, rr := range r {
		rr.Close()
	}
}

// Close closes all the writers, and readers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Close(r []Reader, w Writer) {
	cWriter(w)
	cReader(r)
}
