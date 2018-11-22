// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	type TestStruct struct {
		input       *string
		description string
		expected    bool
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    true,
			description: "is true on nil",
			input:       nil,
		},
		TestStruct{
			input:       &[]string{"string"}[0],
			description: "it's false on non-empty string",
			expected:    false,
		},
		TestStruct{
			input:       &[]string{""}[0],
			description: "it's true on \"\"",
			expected:    true,
		},
	} {
		args := Args{
			"k": ts.input,
		}

		actual := args.IsEmpty("k")
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}
