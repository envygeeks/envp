// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/envygeeks/envp/args"
	"github.com/stretchr/testify/assert"
)

func TestNewFlags(t *testing.T) {
	a := NewFlags()
	assert.NotEmpty(t, a)
}

func TestFlags_Parse(t *testing.T) {
	e := args.Args{}
	NewFlags().Parse()
	a := NewFlags().Parse()
	assert.IsType(t, e, a)
}
