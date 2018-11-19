// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplateString(t *testing.T) {
	tt := template.New("envp")
	tt.Parse("{{ define \"hello\" }}world{{ end }}")
	h := New(tt)

	for _, ttt := range []ts{
		ts{
			e: "world",
			d: "it's world",
			a: "hello",
		},
	} {
		a := h.TemplateString(ttt.a.(string))
		assert.Equal(t, ttt.e, a, ttt.d)
	}
}

func TestTrimmedTemplate(t *testing.T) {
	td := "{{ define \"hello\" }}%s{{ end }}"
	for _, ttt := range []ts{
		ts{
			e: "world",
			d: "it works on simple strings",
			a: "world",
		},
		ts{
			a: "\n\nworld",
			d: "it works on strings with edged space",
			e: "world",
		},
		ts{
			a: "hello\n   \nworld",
			d: "it trims lines with nothing but space",
			e: "hello\n\nworld",
		},
	} {
		tt := template.New("envp")
		tt.Parse(fmt.Sprintf(td, ttt.a.(string)))
		h := New(tt)

		a := h.TrimmedTemplate("hello")
		assert.Equal(t, ttt.e, a, ttt.d)
	}
}

func TestIndentedTemplate(t *testing.T) {
	td := "{{ define \"hello\" }}%s{{ end }}"
	for _, ttt := range []ts{
		ts{
			e: "hello\n  big\nworld",
			a: "\n\n\thello\n\t  big\n\tworld",
			d: "it reindents",
		},
	} {
		tt := template.New("envp")
		tt.Parse(fmt.Sprintf(td, ttt.a.(string)))
		h := New(tt)

		a := h.IndentedTemplate("hello")
		assert.Equal(t, ttt.e, a, ttt.d)
	}
}

func TestTemplateExists(t *testing.T) {
	tt := template.New("envp")
	tt.Parse("{{ define \"hello\" }}world{{ end }}")
	h := New(tt)

	for _, ttt := range []ts{
		ts{
			a: "envp",
			d: "it's true for itself",
			e: true,
		},
		ts{
			a: "hello",
			d: "it's true for define",
			e: true,
		},
		ts{
			e: false,
			d: "it's false if blank",
			a: "",
		},
		ts{
			a: "unknown",
			d: "it's false if it doesn't exist",
			e: false,
		},
	} {
		a := h.TemplateExists(ttt.a.(string))
		assert.Equal(t, ttt.e, a, ttt.d)
	}
}
