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
