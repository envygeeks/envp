// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package flags

import (
	"testing"

	"github.com/envygeeks/envp/flags/args"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	actual := Parse()
	expected := args.Args{}
	assert.IsType(t, expected,
		actual)
}
