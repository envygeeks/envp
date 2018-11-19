// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestSpace(t *testing.T) {
	h := New(template.New("envp"))
	for _, tt := range []ts{
		ts{
			e: " 1",
			d: "it works for simple strings",
			a: "1",
		},
		ts{
			e: " 1",
			d: "it works for strings with space",
			a: "   1 ",
		},
		ts{
			a: "1\n",
			d: "it works with newlines",
			e: " 1",
		},
	} {
		a := h.Space(tt.a.(string), 1)
		assert.Equal(t, tt.e, a, tt.d)
	}
}

func TestReindent(t *testing.T) {
	h := New(template.New("envp"))
	s, e := "\n\n\t1\n\t  2\n\t3", "1\n  2\n3"
	a := h.Reindent(s)

	assert.Equal(t, e, a)
}

func TestTrimEmpty(t *testing.T) {
	h := New(template.New("envp"))
	s, e := "1\n        \n2", "1\n\n2"
	a := h.TrimEmpty(s)

	assert.Equal(t, e, a)
}

func TestTrimEdges(t *testing.T) {
	h := New(template.New("envp"))
	s, e := "\n\n1\n2\n\n", "1\n2"
	a := h.TrimEdges(s)

	assert.Equal(t, e, a)
}
