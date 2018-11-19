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
	h := New(template.New("env"))
	for _, tt := range []ts{
		ts{
			a: "HOME",
			d: "it's true if it exists",
			e: true,
		},
		ts{
			a: "UNKNOWN",
			d: "it's false if it doesn't exist",
			e: false,
		},
		ts{
			a: "BLANK",
			d: "It's true if blank",
			e: true,
		},
	} {
		a := h.EnvExists(tt.a.(string))
		assert.Equal(t, tt.e, a, tt.d)
	}
}

func TestEnv(t *testing.T) {
	os.Setenv("BLANK", "")
	h := New(template.New("env"))
	for _, tt := range []ts{
		ts{
			a: "HOME",
			d: "it's a string if it exists",
			e: "",
		},
		ts{
			a: "UNKNOWN",
			d: "it's a string if it doesn't exist",
			e: "",
		},
		ts{
			a: "BLANK",
			d: "it's a string if blank",
			e: "",
		},
	} {
		a := h.Env(tt.a.(string))
		assert.IsType(t, tt.e, a, tt.d)
	}
}

func TestTemplate__boolEnv(t *testing.T) {
	os.Setenv("TRUE_1", "1")
	os.Setenv("TRUE_TRUE", "true")
	os.Setenv("FALSE_FALSE", "false")
	os.Setenv("FALSE_0", "0")
	os.Setenv("BLANK", "")

	h := New(template.New("envp"))
	for _, tt := range []ts{
		ts{
			a: "FALSE_0",
			d: "it's false if it's 0",
			e: false,
		},
		ts{
			a: "UNKNOWN",
			d: "it's false if it doesn't exist",
			e: false,
		},
		ts{
			a: "TRUE_TRUE",
			d: "it's true if it's true",
			e: true,
		},
		ts{
			a: "HOME",
			d: "it's false if it exists, and isn't true/1",
			e: false,
		},
		ts{
			a: "FALSE_FALSE",
			d: "it's false if it's false",
			e: false,
		},
		ts{
			a: "BLANK",
			d: "it's false if it's blank",
			e: false,
		},
		ts{
			a: "TRUE_1",
			d: "it's true if it's 1",
			e: true,
		},
	} {
		a := h.BoolEnv(tt.a.(string))
		assert.Equal(t, tt.e, a, tt.d)
	}
}
