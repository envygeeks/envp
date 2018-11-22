// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestEnvExists(t *testing.T) {
	os.Setenv("BLANK", "")
	type TestStruct struct {
		expected    bool
		description string
		key         string
	}

	h := New(template.New("env"))
	for _, ts := range []TestStruct{
		TestStruct{
			key:         "HOME",
			description: "it's true if it exists",
			expected:    true,
		},
		TestStruct{
			key:         "UNKNOWN",
			description: "it's false if it doesn't exist",
			expected:    false,
		},
		TestStruct{
			key:         "BLANK",
			description: "It's true if blank",
			expected:    true,
		},
	} {
		actual := h.EnvExists(ts.key)
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("BLANK", "")
	type TestStruct struct {
		expected    string
		description string
		key         string
	}

	h := New(template.New("env"))
	for _, ts := range []TestStruct{
		TestStruct{
			key:         "HOME",
			description: "it's a string if it exists",
			expected:    "",
		},
		TestStruct{
			key:         "UNKNOWN",
			description: "it's a string if it doesn't exist",
			expected:    "",
		},
		TestStruct{
			key:         "BLANK",
			description: "it's a string if blank",
			expected:    "",
		},
	} {
		actual := h.Env(ts.key)
		assert.IsType(t, ts.expected, actual,
			ts.description)
	}
}

func TestTemplate__boolEnv(t *testing.T) {
	os.Setenv("TRUE_1", "1")
	os.Setenv("TRUE_TRUE", "true")
	os.Setenv("FALSE_FALSE", "false")
	os.Setenv("FALSE_0", "0")
	os.Setenv("BLANK", "")

	type TestStruct struct {
		expected    bool
		description string
		key         string
	}

	h := New(template.New("envp"))
	for _, ts := range []TestStruct{
		TestStruct{
			expected:    false,
			description: "it's false if it's 0",
			key:         "FALSE_0",
		},
		TestStruct{
			expected:    false,
			description: "it's false if it doesn't exist",
			key:         "UNKNOWN",
		},
		TestStruct{
			key:         "TRUE_TRUE",
			description: "it's true if it's true",
			expected:    true,
		},
		TestStruct{
			expected:    false,
			description: "it's false if it exists, and isn't true/1",
			key:         "HOME",
		},
		TestStruct{
			key:         "FALSE_FALSE",
			description: "it's false if it's false",
			expected:    false,
		},
		TestStruct{
			expected:    false,
			description: "it's false if it's blank",
			key:         "BLANK",
		},
		TestStruct{
			expected:    true,
			description: "it's true if it's 1",
			key:         "TRUE_1",
		},
	} {
		actual := h.BoolEnv(ts.key)
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestSpace(t *testing.T) {
	h := New(template.New("envp"))
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    " 1",
			description: "it works for simple strings",
			input:       "1",
		},
		TestStruct{
			expected:    " 1",
			description: "it works for strings with space",
			input:       "   1 ",
		},
		TestStruct{
			input:       "1\n",
			description: "it works with newlines",
			expected:    " 1",
		},
	} {
		actual := h.Space(ts.input, 1)
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestReindent(t *testing.T) {
	h := New(template.New("envp"))
	input, expected := "\n\n\t1\n\t  2\n\t3", "1\n  2\n3"
	actual := h.Reindent(input)
	assert.Equal(t, expected,
		actual)
}

func TestTrimEmpty(t *testing.T) {
	h := New(template.New("envp"))
	input, expected := "1\n        \n2", "1\n\n2"
	actual := h.TrimEmpty(input)

	assert.Equal(t, expected,
		actual)
}

func TestTrimEdges(t *testing.T) {
	h := New(template.New("envp"))
	input, expected := "\n\n1\n2\n\n", "1\n2"
	actual := h.TrimEdges(input)

	assert.Equal(t, expected,
		actual)
}

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
