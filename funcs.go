// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

const (
	rrRE = `(?m)^[ \t]{%d}`
	teRE = `(?m)\A[ \t]*$[\r\n]*|[\r\n]+[ \t]*\z`
	rdRE = `(?m)^[ \t]*`
)

// _templateString allows you to pull a template
// and output it as a string into a func or whatever.
func (t *Template) _templateString(s string) string {
	if tt := t.template.Lookup(s); tt != nil {
		var ss strings.Builder
		if err := tt.Execute(&ss, t); err != nil {
			log.Fatalln(err)
		}

		return ss.String()
	}

	// Bad template given.
	log.Fatalf("Unable to find %s", s)
	return ""
}

// _reindentedTemplate reindents, and outs a string
func (t *Template) _reindentedTemplate(s string) string {
	s = t._templateString(s)
	s = t._reindent(s)
	return s
}

// _reIndent takes a string, and strips the
// indention to the edge, like Rails #strip_heredoc
// or Ruby std <<~, it also strips blank lines on
// the top and on the bottom, for swift alignment
func (t *Template) _reindent(s string) string {
	s, indent := t._trimEdges(s), -1
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

// _trimEmptyLines trims empty lines on the
// top and on the bottom of a string so that you
// can do something close to Ruby's <<~
func (t *Template) _trimEdges(s string) string {
	re := regexp.MustCompile(teRE)
	return re.ReplaceAllString(
		s, "")
}

// _space adds a space to the beginning of
// a string this way you can {{- -}} compress
// lines, and still have a space
func (t *Template) _space(s string, n int) string {
	s = strings.TrimSpace(s)
	space := strings.Repeat(" ", n)
	return space + s
}

// _templateExists allows you to check if a
// template exists inside of the templates, this also
// works for context based {{ define "name" }}.
func (t *Template) _templateExists(s string) bool {
	log.Debugf("looking for template %s", s)
	if tt := t.template.Lookup(s); tt != nil {
		return true
	}

	return false
}

// _envExists allows you to check if a var exists
// in your current environment, we do not alter it so
// make sure you use FULLCAP if necessary.
func (t *Template) _envExists(s string) bool {
	_, ok := os.LookupEnv(s)
	log.Debugf("checked if env %s exists", s)
	return ok
}

// _env allows you to pull out a string env var
func (t *Template) _env(s string) string {
	if v, ok := os.LookupEnv(s); ok {
		return v
	}

	return ""
}

// _boolEnv allows you to pull out a env var as a
// bool, following the same rules as strconv.ParseBool
// where 1, true are true, and all else is false
func (t *Template) _boolEnv(s string) bool {
	if v, ok := os.LookupEnv(s); ok {
		vv, err := strconv.ParseBool(v)
		if err != nil {
			log.Warn(err)
			return false
		}

		return vv
	}

	return false
}

// addFuncs attaches the funcs to the template
func (t *Template) addFuncs() *Template {
	t.template.Funcs(template.FuncMap{
		"split":              strings.Split,
		"boolEnv":            t._boolEnv,
		"reindent":           t._reindent,
		"trimEdges":          t._trimEdges,
		"templateString":     t._templateString,
		"reindentedTemplate": t._reindentedTemplate,
		"templateExists":     t._templateExists,
		"envExists":          t._envExists,
		"trim":               strings.Trim,
		"env":                t._env,
	})

	return t
}
