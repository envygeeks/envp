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
	upstream "text/template"

	"github.com/envygeeks/envp/template/helpers"
	"github.com/sirupsen/logrus"
)

// Template provides a context wrapper
// for all of our internal functions, for
// the template, and stuff for you.
type Template struct {
	*upstream.Template

	use   string
	debug bool
}

// New creates a new template, and logs it for
// the entire world to know if they really need to
// know what's going on for debugging purposes.
func New() *Template {
	upstream := upstream.New("envp")
	template := &Template{
		Template: upstream,
	}

	helpers.New(upstream)
	return template
}

// Use tells us to use this specific template
func (t *Template) Use(f Reader) {
	t.use = filepath.Base(f.Name())
}

// ParseFiles parses all your readers
func (t *Template) ParseFiles(readers []Reader) []*upstream.Template {
	var templates []*upstream.Template

	for _, v := range readers {
		if reader, ok := v.(Reader); ok {
			templates = append(templates, t.ParseFile(reader))
		}
	}

	return templates
}

// ParseFile parses a reader into a template
func (t *Template) ParseFile(reader Reader) *upstream.Template {
	logrus.Debugf("attempting to add & parse %s", reader.Name())
	template := t.New(filepath.Base(reader.Name()))
	if byte, err := ioutil.ReadAll(reader); err != nil {
		logrus.Fatalln(err)
	} else {
		if _, err := template.Parse(string(byte)); err != nil {
			logrus.Fatalln(err)
		}
	}

	return template
}

// Compile runs exec on the template.
// Before you hit this stage you should really be
// running Load(), and Parse() to get ready.
func (t *Template) Compile() []byte {
	var template *upstream.Template

	if t.use != "" {
		logrus.Debugf("using requested %s", t.use)
		template = t.Lookup(t.use)
		if template == nil {
			logrus.Fatalf("unable to find %s", t.use)
		}
	} else {
		templates := t.Templates()
		for _, v := range templates {
			if v.Name() == "base.gohtml" || v.Name() == "root.gohtml" {
				template = v
				break
			}
		}

		if template == nil {
			template = templates[0]
			if template == nil {
				logrus.Fatalln("no template found")
			}
		}
	}

	buf := &bytes.Buffer{}
	logrus.Debugf("executing %s", template.Name())
	if err := template.Execute(buf, ""); err != nil {
		logrus.Fatalln(err)
	}

	return buf.Bytes()
}

// Writer interface
type Writer interface {
	Name() string
	Close() error
	io.Writer
}

// Write writes to stdout, or a file.
func (t *Template) Write(b []byte, w Writer) int {
	oint, err := w.Write(b)
	if err != nil {
		logrus.Fatalln(err)
	}

	return oint
}

func writer(file string) *os.File {
	var mode os.FileMode
	if file == "" {
		logrus.Infoln("using stdout")
		return os.Stdout
	}

	file, err := filepath.Abs(file)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Debugf("opening a writer to %s", file)
	mode, op := 0644, os.O_CREATE|os.O_WRONLY
	writer, err := os.OpenFile(file, op, mode)
	if err != nil {
		logrus.Fatalln(err)
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
	logrus.Debugf("opening a reader to %s", file)
	reader, err := os.Open(file)
	if err != nil {
		logrus.Fatalln(err)
	}

	return reader
}

var (
	globName = "*.gohtml"
)

func readers(files []string) []Reader {
	var readers []Reader

	for _, file := range files {
		abspath, err := filepath.Abs(file)
		if err != nil {
			logrus.Fatalln(err)
		}

		finfo, err := os.Stat(abspath)
		if err != nil {
			logrus.Fatalln(err)
		} else {
			if !finfo.IsDir() {
				readers = append(readers, reader(file))
				continue
			}
		}

		path := filepath.Join(file, globName)
		logrus.Infof("looking for %s in %s", globName, file)
		all, err := filepath.Glob(path)
		if err != nil {
			logrus.Fatalln(err)
		}

		for _, gfile := range all {
			readers = append(readers, reader(gfile))
		}

	}

	return readers
}

// Open opens all the readers, and writers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Open(rs []string, w string) ([]Reader, Writer) {
	return readers(rs), writer(w)
}

/**
 */
func closeWriter(writer Writer) bool {
	if writer != os.Stdout {
		if err := writer.Close(); err == nil {
			return true
		}
	}

	return false
}

/**
 */
func closeReader(readers []Reader) bool {
	var closed bool

	for _, reader := range readers {
		closed = true
		if err := reader.Close(); err != nil {
			closed = false
		}
	}

	return closed
}

// Close closes all the writers, and readers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Close(r []Reader, w Writer) {
	closeWriter(w)
	closeReader(r)
}
