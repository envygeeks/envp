// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	for _, v := range [][3]interface{}{
		{nil, true, "is true on nil"},
		{"string", false, "it false on non-empty string"},
		{"", true, "it true on \"\""},
	} {
		args := Args{"k": v[0]}
		assert.Equal(t, args.IsEmpty("k"), v[1], v[2])
	}
}
