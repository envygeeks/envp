// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
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
