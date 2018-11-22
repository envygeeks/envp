// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/envygeeks/envp/logger"
)

// Helpers provides the struct that holds
// the template, for functions that require such
// things, we use store it.
type Helpers struct {
	template *template.Template
}

// New creates a new Funcs
func New(t *template.Template) *Helpers {
	return (&Helpers{
		template: t,
	}).register()
}

// Register registers the funcs
func (h *Helpers) register() *Helpers {
	logger.Println("registering all the helpers")
	h.template.Funcs(template.FuncMap{
		"trim":             strings.Trim,
		"split":            strings.Split,
		"boolEnv":          h.BoolEnv,
		"envExists":        h.EnvExists,
		"env":              h.Env,
		"reindent":         h.Reindent,
		"trimEmpty":        h.TrimEmpty,
		"trimEdges":        h.TrimEdges,
		"space":            h.Space,
		"templateString":   h.TemplateString,
		"trimmedTemplate":  h.TrimmedTemplate,
		"indentedTemplate": h.IndentedTemplate,
		"templateExists":   h.TemplateExists,
	})

	return h
}

// EnvExists allows you to check if a var exists
// in your current environment, we do not alter it so
// make sure you use FULLCAP if necessary.
func (h *Helpers) EnvExists(s string) bool {
	s = strings.ToUpper(s)
	if _, ok := os.LookupEnv(s); ok {
		return true
	}

	return false
}

// Env allows you to pull out a string env var
func (h *Helpers) Env(s string) string {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		return v
	}

	return ""
}

// BoolEnv allows you to pull out a env var as a
// bool, following the same rules as strconv.ParseBool
// where 1, true are true, and all else is false
func (h *Helpers) BoolEnv(s string) bool {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		vv, err := strconv.ParseBool(v)
		if err != nil {
			logger.Println(err)
			return false
		}

		return vv
	}

	return false
}

// Space adds a space to the beginning of
// a string this way you can {{- -}} compress
// lines, and still have a space
func (h *Helpers) Space(s string, n int) string {
	s = strings.TrimSpace(s)
	space := strings.Repeat(" ", n)
	return space + s
}

const (
	elRE = `(?m)^[ \t]+$`
	rrRE = `(?m)^[ \t]{%d}`
	teRE = `(?m)\A[ \t]*$[\r\n]*|[\r\n]+[ \t]*\z`
	rdRE = `(?m)^[ \t]+`
)

// Reindent takes a string, and strips the
// indention to the edge, like Rails #strip_heredoc
// or Ruby std <<~, it also strips blank lines on
// the top and on the bottom, for swift alignment
func (h *Helpers) Reindent(s string) string {
	s, indent := h.TrimEdges(s), -1
	re := regexp.MustCompile(rdRE)
	for _, v := range re.FindAllString(s, -1) {
		if l := len(v); indent == -1 || l < indent {
			indent = l
		}
	}

	if indent > -1 {
		re := regexp.MustCompile(fmt.Sprintf(rrRE, indent))
		s = re.ReplaceAllString(s, "")
	}

	return s
}

// TrimEmpty strips empty lines of any
// space that is pushed through by templating
// allowing you to work with simply \n
func (h *Helpers) TrimEmpty(s string) string {
	re := regexp.MustCompile(elRE)
	return re.ReplaceAllString(s, "")
}

// TrimEdges trims empty lines on the
// top and on the bottom of a string so that you
// can do something close to Ruby's <<~
func (h *Helpers) TrimEdges(s string) string {
	re := regexp.MustCompile(teRE)
	return re.ReplaceAllString(s, "")
}

// TemplateString allows you to pull a template
// and output it as a string into a func or whatever
// pleases you, this is useful for storing bits of
// a template into a variable for later.
func (h *Helpers) TemplateString(s string) string {
	if tt := h.template.Lookup(s); tt != nil {
		var ss strings.Builder
		if err := tt.Execute(&ss, h.template); err != nil {
			logger.Fatalln(err)
		}

		return ss.String()
	}

	// Bad template given.
	logger.Fatalf("Unable to find %s", s)
	return ""
}

// TrimmedTemplate trims the TemplateString, this
// should generally give you a cleaner template output
// that is more suited to configuration files.
func (h *Helpers) TrimmedTemplate(s string) string {
	return h.TrimEmpty(h.TrimEdges(h.TemplateString(s)))
}

// IndentedTemplate reindents, and outs a string
// template, this also runs TrimmedTemplate so that
// you end up with a clean, and indented template
// for human readable configuration files.
func (h *Helpers) IndentedTemplate(s string) string {
	return h.Reindent(h.TrimmedTemplate(s))
}

// TemplateExists allows you to check if a
// template exists inside of the templates, this also
// works for context based {{ define "name" }}.
func (h *Helpers) TemplateExists(s string) bool {
	logger.Printf("looking for template %s", s)
	if tt := h.template.Lookup(s); tt != nil {
		return true
	}

	return false
}
