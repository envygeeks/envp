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

// Helpers holds the template
// We use it for some template funcs
// Some don't use it at all.
type Helpers struct {
	template *template.Template
}

// EnvExists allows you to check if a var exists
func (h *Helpers) EnvExists(s string) bool {
	s = strings.ToUpper(s)
	if _, ok := os.LookupEnv(s); ok {
		return true
	}

	return false
}

// Env allows you to pull out a string var
func (h *Helpers) Env(s string) string {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		return v
	}

	return ""
}

// BoolEnv allows you to pull out a var as bool
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

// AddSpace adds a space to the beginning of a string
func (h *Helpers) AddSpace(s string, n int) string {
	s = strings.TrimSpace(s)
	space := strings.Repeat(" ", n)
	return space + s
}

const (
	indentRegex     = `(?m)^[ \t]{%d}`
	stripEdgesRegex = `(?m)\A[ \t]*$[\r\n]*|[\r\n]+[ \t]*\z`
	stripEmptyRegex = `(?m)^[ \t]+$`
	preindentRegex  = `(?m)^[ \t]+`
)

func indent(h *Helpers, s string, size int) string {
	re, s, cSize := regexp.MustCompile(preindentRegex), h.Strip(s), -1
	for _, v := range re.FindAllString(s, -1) {
		if l := len(v); cSize == -1 || l < cSize {
			cSize = l
			continue
		}
	}

	if cSize > -1 {
		rr := fmt.Sprintf(indentRegex, cSize)
		re := regexp.MustCompile(rr)
		if size > -1 {
			i := strings.Repeat(" ", size)
			return re.ReplaceAllString(
				s, i)
		}

		return re.ReplaceAllString(s, "")
	}

	return s
}

// FixIndentation strips indentation to the edge
func (h *Helpers) FixIndentation(s string) string {
	o := indent(h, s, -1)
	return o
}

// Indent strips each lines indentation to the
// edge, and then indents each line to the size you set
func (h *Helpers) Indent(s string, size uint) string {
	o := indent(h, s, int(size))
	return o
}

// Strip removes empty lines around a string
func (h *Helpers) Strip(s string) string {
	re1 := regexp.MustCompile(stripEmptyRegex)
	re2 := regexp.MustCompile(stripEdgesRegex)
	s = re1.ReplaceAllString(s, "")
	s = re2.ReplaceAllString(s, "")
	return s
}

// TemplateString pulls a template as a string
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

// StrippedTemplate trims empty lines, and edges.
func (h *Helpers) StrippedTemplate(s string) string {
	s = h.TemplateString(s)
	s = h.Strip(s)
	return s
}

// FixIndentedTemplate strips the indentation to the edge
func (h *Helpers) FixIndentedTemplate(s string) string {
	s = h.TemplateString(s)
	s = h.FixIndentation(s)
	return s
}

// TemplateExists checks if a template exists
func (h *Helpers) TemplateExists(s string) bool {
	logger.Printf("looking for template %s", s)
	if tt := h.template.Lookup(s); tt != nil {
		return true
	}

	return false
}

func New(t *template.Template) *Helpers {
	return (&Helpers{
		template: t,
	}).RegisterFuncs()
}

// RegisterFuncs registers the funcs
func (h *Helpers) RegisterFuncs() *Helpers {
	logger.Println("registering all the helpers")
	h.template.Funcs(template.FuncMap{
		"split":               strings.Split,
		"chomp":               strings.Trim,
		"indent":              h.Indent,
		"addSpace":            h.AddSpace,
		"templateString":      h.TemplateString,
		"strippedTemplate":    h.StrippedTemplate,
		"fixIndentedTemplate": h.FixIndentedTemplate,
		"templateExists":      h.TemplateExists,
		"fixIndentation":      h.FixIndentation,
		"envExists":           h.EnvExists,
		"boolEnv":             h.BoolEnv,
		"strip":               h.Strip,
		"env":                 h.Env,
	})

	return h
}
