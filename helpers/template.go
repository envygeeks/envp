// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

// TemplateString allows you to pull a template
// and output it as a string into a func or whatever
// pleases you, this is useful for storing bits of
// a template into a variable for later.
func (h *Helpers) TemplateString(s string) string {
	if tt := h.template.Lookup(s); tt != nil {
		var ss strings.Builder
		if err := tt.Execute(&ss, h.template); err != nil {
			log.Fatalln(err)
		}

		return ss.String()
	}

	// Bad template given.
	log.Fatalf("Unable to find %s", s)
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
	log.Debugf("looking for template %s", s)
	if tt := h.template.Lookup(s); tt != nil {
		return true
	}

	return false
}
