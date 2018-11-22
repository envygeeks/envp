// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
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
