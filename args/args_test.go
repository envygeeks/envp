// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ts struct {
	a interface{}
	e interface{}
	d string
}

func TestNew(t *testing.T) {
	a, e := New(), Args{}
	assert.IsType(t, e, a)
}
