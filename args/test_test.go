// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	for _, tt := range []ts{
		ts{
			e: true,
			d: "is true on nil",
			a: nil,
		},
		ts{
			a: "string",
			d: "it's false on non-empty string",
			e: false,
		},
		ts{
			e: true,
			d: "it's true on \"\"",
			a: "",
		},
	} {
		args := Args{
			"k": tt.a,
		}

		a := args.IsEmpty("k")
		assert.Equal(t, tt.e, a, tt.d)
	}
}
