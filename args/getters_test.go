// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	type TestStruct struct {
		input       *bool
		description string
		expected    bool
	}

	for _, ts := range []TestStruct{
		TestStruct{
			input:       &[]bool{false}[0],
			description: "it's false when false",
			expected:    false,
		},
		TestStruct{
			input:       &[]bool{true}[0],
			description: "it's true when true",
			expected:    true,
		},
	} {
		args := Args{
			"k": ts.input,
		}

		actual := args.Bool("k")
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestString(t *testing.T) {
	type TestStruct struct {
		input       *string
		description string
		expected    string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			input:       &[]string{""}[0],
			description: "it returns a string",
			expected:    "",
		},
	} {
		args := Args{
			"k": ts.input,
		}

		actual := args.String("k")
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}
