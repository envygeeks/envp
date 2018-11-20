// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	e := os.Stdout
	a := writer("", true)
	assert.IsType(t, e, a)
}
