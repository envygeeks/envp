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

	helpers := New(template.New("env"))
	for _, test := range []TestStruct{
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
		actual := helpers.EnvExists(test.key)
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("BLANK", "")
	type TestStruct struct {
		expected    string
		description string
		key         string
	}

	helpers := New(template.New("env"))
	for _, test := range []TestStruct{
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
		actual := helpers.Env(test.key)
		assert.IsType(t, test.expected, actual,
			test.description)
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

	helpers := New(template.New("envp"))
	for _, test := range []TestStruct{
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
		actual := helpers.BoolEnv(test.key)
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestAddSpace(t *testing.T) {
	helpers := New(template.New("envp"))
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, test := range []TestStruct{
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
		actual := helpers.AddSpace(test.input, 1)
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestFixIndentation(t *testing.T) {
	helpers := New(template.New("envp"))
	input, expected := "\n\n\t1\n\t  2\n\t3", "1\n  2\n3"
	actual := helpers.FixIndentation(input)
	assert.Equal(t, expected,
		actual)
}
func TestIndent(t *testing.T) {
	helpers := New(template.New("envp"))
	input, expected := "\n\n\t1\n\t  2\n\t3", "  1\n    2\n  3"
	actual := helpers.Indent(input, 2)
	assert.Equal(t, expected,
		actual)
}

func TestStrip(t *testing.T) {
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, test := range []TestStruct{
		TestStruct{
			expected:    "1\n\n2",
			description: "it works on simple strings",
			input:       "1\n        \n2",
		},
		TestStruct{
			input:       "\n\n1\n2\n\n",
			description: "it works on strings with edged space",
			expected:    "1\n2",
		},
	} {
		actual := New(template.New("envp")).Strip(test.input)
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestTemplateString(t *testing.T) {
	template := template.New("envp")
	template.Parse("{{ define \"hello\" }}world{{ end }}")
	helpers := New(template)

	type TestStruct struct {
		expected    string
		description string
		key         string
	}

	for _, test := range []TestStruct{
		TestStruct{
			expected:    "world",
			description: "it's world",
			key:         "hello",
		},
	} {
		actual := helpers.TemplateString(test.key)
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestStrippedTemplate(t *testing.T) {
	tpldef := "{{ define \"hello\" }}%s{{ end }}"
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, test := range []TestStruct{
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
		template := template.New("envp")
		template.Parse(fmt.Sprintf(tpldef, test.input))
		helpers := New(template)

		actual := helpers.StrippedTemplate("hello")
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestFixIndentedTemplate(t *testing.T) {
	tpldef := "{{ define \"hello\" }}%s{{ end }}"
	type TestStruct struct {
		expected    string
		description string
		input       string
	}

	for _, test := range []TestStruct{
		TestStruct{
			expected:    "hello\n  big\nworld",
			description: "it reindents the template properly",
			input:       "\n\n\thello\n\t  big\n\tworld",
		},
	} {
		template := template.New("envp")
		template.Parse(fmt.Sprintf(tpldef, test.input))
		h := New(template)

		actual := h.FixIndentedTemplate("hello")
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestTemplateExists(t *testing.T) {
	template := template.New("envp")
	template.Parse("{{ define \"hello\" }}world{{ end }}")
	helpers := New(template)

	type TestStruct struct {
		expected    bool
		description string
		key         string
	}

	for _, test := range []TestStruct{
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
		actual := helpers.TemplateExists(test.key)
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestRandomPassword(t *testing.T) {
	helpers := New(template.New("envp"))
	actual := helpers.RandomPassword(12)
	assert.Equal(t, 12, len(actual))
	assert.NotNil(t, actual)
}
