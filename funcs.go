// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"strconv"
	"strings"
	"text/template"
)

type BoolFunc func(s string) bool
type BoolFuncWithError func(s string) (bool, error)
type StringFuncWithError func(s string) (string, error)
type IntFuncWithError func(s string) (int, error)
type StringFunc func(s string) string
type IntFunc func(s string) int

// func_trimStr trims a string for you... obviously
func func_trimStr(t *template.Template) StringFunc {
	return func(s string) string {
		return strings.TrimSpace(s)
	}
}

// func_templateExists allows you to check if a
// template exists inside of the templates, this also
// works for context based {{ define "name" }}.
func func_templateExists(t *template.Template) BoolFunc {
	return func(s string) bool {
		if tt := t.Lookup(s); tt != nil {
			return true
		}

		return false
	}
}

// func_eExists allows you to check if a var exists
// in your current environment, we do not alter it so
// make sure you use FULLCAP if necessary.
func func_eExists(t *template.Template) BoolFunc {
	return func(s string) bool {
		_, ok := os.LookupEnv(s)
		return ok
	}
}

// func_eStr allows you to pull out a string env var
func func_eStr(t *template.Template) StringFunc {
	return func(s string) string {
		if v, ok := os.LookupEnv(s); ok {
			return v
		}

		return ""
	}
}

// func_eBool allows you to pull out a env var as a
// bool, following the same rules as strconv.ParseBool
// where 1, true are true, and all else is false
func func_eBool(t *template.Template) BoolFuncWithError {
	return func(s string) (bool, error) {
		if v, ok := os.LookupEnv(s); ok {
			vv, err := strconv.ParseBool(v)
			if err != nil {
				return false, err
			}

			return vv, nil
		}

		return false, nil
	}
}

func funcs(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"eBool":          func_eBool(t),
		"templateExists": func_templateExists(t),
		"eExists":        func_eExists(t),
		"trimStr":        func_trimStr(t),
		"eStr":           func_eStr(t),
	}
}
