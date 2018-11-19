// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ts struct {
	a interface{}
	e interface{}
	d string
}

func TestIsExist(t *testing.T) {
	for _, tt := range []ts{
		ts{
			e: true,
			d: "it's true when it does",
			a: "/",
		},
		ts{
			a: "/should/not/exist",
			d: "it's false when it doesn't",
			e: false,
		},
	} {
		a := IsExist(tt.a.(string))
		assert.Equal(t, tt.e, a, tt.d)
	}
}
