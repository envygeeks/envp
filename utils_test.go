// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	for _, v := range [][3]interface{}{
		{"etc", "it expands"},
	} {
		actual := strings.HasPrefix(expand(v[0].(string)), "/")
		assert.True(t, actual, v[1])
	}
}

func TestIsExist(t *testing.T) {
	for _, v := range [][3]interface{}{
		{true, "/etc", "it's true when it does"},
		{false, "/should/not/exist", "it's false when it doesn't"},
	} {
		actual := isExist(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}
