package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// Template provides a context wrapper
// for all of our internal functions, for
// the template, and stuff for you.
type Template struct {
	file      string
	output    string
	template  *template.Template
	templates []string
	stdout    bool
	debug     bool
	glob      bool
}

// NewTemplate creates a new template, and logs it
// for the entire world to know if they really need to
// know what's going on for debugging purposes.
func NewTemplate(a *Args) *Template {
	log.Debug("creating new envp template")
	return (&Template{
		debug:    a.Bool("debug"),
		file:     expand(a.String("file")),
		output:   expand(a.String("output")),
		template: template.New("envp"),
		stdout:   a.Bool("stdout"),
		glob:     a.Bool("glob"),
	}).Verify().addFuncs()
}

// Verify verifies the file exists.
func (t *Template) Verify() *Template {
	if !isExist(t.file) {
		log.Fatalf("%s doesn't exist", t.file)
	}

	return t
}

// Load loads the templates.
// If you add new templates, you can rerun
// Load() and Parse(), or Run()
func (t *Template) Load() {
	if t.glob {
		var err error
		log.Infof("searching for *.gohtml in %s", t.file)
		p := filepath.Join(t.file, "*.gohtml")
		t.templates, err = filepath.Glob(p)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		t.templates = []string{
			t.file,
		}
	}
}

// Parse parses the templates.
func (t *Template) Parse() {
	log.Printf("attempting to parse %+v", t.templates)
	if _, err := t.template.ParseFiles(t.templates...); err != nil {
		log.Warnf("unable to parse %+v", t.templates)
		log.Fatalln(err)
	}
}

// Exec runs exec on the template.
// Before you hit this stage you should really be
// running Load(), and Parse() to get ready.
func (t *Template) Exec() *strings.Builder {
	var tt *template.Template

	if len(t.templates) == 1 {
		name := filepath.Base(t.templates[0])
		log.Infof("looking for %s", name)
		if tt = t.template.Lookup(name); tt == nil {
			log.Fatalf("no template \"%s\"", name)
		}
	} else {
		for _, v := range []string{"base.gohtml", "root.gohtml"} {
			log.Infof("looking for %s", v)
			tt = t.template.Lookup(v)
			if tt != nil {
				break
			}
		}

		if tt == nil {
			log.Fatal("no base, or root")
		}
	}

	var s strings.Builder
	log.Infof("executing %s", tt.Name())
	if err := tt.Execute(&s, ""); err != nil {
		log.Fatal(err)
	}

	return &s
}

// Write writes to stdout, or a file.
func (t *Template) Write(s *strings.Builder) {
	if !t.stdout {
		d := filepath.Dir(t.output)
		log.Infof("writing %s", t.output)
		if err := os.MkdirAll(d, 0644); err != nil {
			log.Fatalln(err)
		} else {
			b := []byte(s.String())
			if err := ioutil.WriteFile(t.output, b, 0644); err != nil {
				log.Fatalln(err)
			}
		}
	} else {
		if t.debug {
			fmt.Print("\n\n")
			fmt.Print("\n\n")
			fmt.Print("\n\n")
		}

		fmt.Println(s.String())
	}
}

// Run runs Load(), Parse()
func (t *Template) Run() {
	t.Load()
	t.Parse()
	s := t.Exec()
	t.Write(s)
}
