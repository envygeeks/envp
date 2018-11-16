// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"strconv"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// func_trimStr trims a string for you... obviously
func (t *Template) _trimStr(s string) string {
	return strings.TrimSpace(s)
}

// func_templateExists allows you to check if a
// template exists inside of the templates, this also
// works for context based {{ define "name" }}.
func (t *Template) _templateExists(s string) bool {
	log.Debugf("looking for template %s", s)
	if tt := t.template.Lookup(s); tt != nil {
		return true
	}

	return false
}

// func_eExists allows you to check if a var exists
// in your current environment, we do not alter it so
// make sure you use FULLCAP if necessary.
func (t *Template) _eExists(s string) bool {
	_, ok := os.LookupEnv(s)
	log.Debugf("checked if env %s exists", s)
	return ok
}

// func_eStr allows you to pull out a string env var
func (t *Template) _eStr(s string) string {
	if v, ok := os.LookupEnv(s); ok {
		return v
	}

	return ""
}

// func_eBool allows you to pull out a env var as a
// bool, following the same rules as strconv.ParseBool
// where 1, true are true, and all else is false
func (t *Template) _eBool(s string) bool {
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

// setupFuncs attaches the funcs to the template
func (t *Template) addFuncs() *Template {
	t.template.Funcs(template.FuncMap{
		"eBool":          t._eBool,
		"templateExists": t._templateExists,
		"eExists":        t._eExists,
		"trimStr":        t._trimStr,
		"eStr":           t._eStr,
	})

	return t
}
