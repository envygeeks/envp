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

	type TestStruct struct {
		expected    string
		description string
		key         string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    "world",
			description: "it's world",
			key:         "hello",
		},
	} {
		actual := h.TemplateString(ts.key)
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestTrimmedTemplate(t *testing.T) {
	templateDefinition := "{{ define \"hello\" }}%s{{ end }}"
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    "world",
			description: "it works on simple strings",
			input:       "world",
		},
		TestStruct{
			input:       "\n\nworld",
			description: "it works on strings with edged space",
			expected:    "world",
		},
		TestStruct{
			input:       "hello\n   \nworld",
			description: "it trims lines with nothing but space",
			expected:    "hello\n\nworld",
		},
	} {
		tt := template.New("envp")
		tt.Parse(fmt.Sprintf(templateDefinition, ts.input))
		h := New(tt)

		actual := h.TrimmedTemplate("hello")
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestIndentedTemplate(t *testing.T) {
	templateDefinition := "{{ define \"hello\" }}%s{{ end }}"
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    "hello\n  big\nworld",
			description: "it reindents the template properly",
			input:       "\n\n\thello\n\t  big\n\tworld",
		},
	} {
		tt := template.New("envp")
		tt.Parse(fmt.Sprintf(templateDefinition, ts.input))
		h := New(tt)

		actual := h.IndentedTemplate("hello")
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestTemplateExists(t *testing.T) {
	tt := template.New("envp")
	tt.Parse("{{ define \"hello\" }}world{{ end }}")
	h := New(tt)

	type TestStruct struct {
		expected    bool
		description string
		key         string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    true,
			description: "it's true for itself",
			key:         "envp",
		},
		TestStruct{
			expected:    true,
			description: "it's true for define",
			key:         "hello",
		},
		TestStruct{
			expected:    false,
			description: "it's false if blank",
			key:         "",
		},
		TestStruct{
			expected:    false,
			description: "it's false if it doesn't exist",
			key:         "unknown",
		},
	} {
		actual := h.TemplateExists(ts.key)
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}
