// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	for _, tt := range []ts{
		ts{
			a: &[]bool{false}[0],
			d: "it's false when false",
			e: false,
		},
		ts{
			a: &[]bool{true}[0],
			d: "it's true when true",
			e: true,
		},
	} {
		args := Args{
			"k": tt.a,
		}

		a := args.Bool("k")
		assert.IsType(t, tt.e, a, tt.d)
	}
}

func TestString(t *testing.T) {
	for _, tt := range []ts{
		ts{
			a: &[]string{""}[0],
			d: "it returns a string",
			e: "",
		},
	} {
		args := Args{
			"k": tt.a,
		}

		a := args.String("k")
		assert.IsType(t, tt.e, a, tt.d)
	}
}
