// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	elRE = `(?m)^[ \t]+$`
	rrRE = `(?m)^[ \t]{%d}`
	teRE = `(?m)\A[ \t]*$[\r\n]*|[\r\n]+[ \t]*\z`
	rdRE = `(?m)^[ \t]+`
)

// Space adds a space to the beginning of
// a string this way you can {{- -}} compress
// lines, and still have a space
func (h *Helpers) Space(s string, n int) string {
	s = strings.TrimSpace(s)
	space := strings.Repeat(" ", n)
	return space + s
}

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
