// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFlags(t *testing.T) {
	assert.NotEmpty(t, NewFlags())
}

func TestFlags_Parse(t *testing.T) {
	a := NewFlags().Parse()
	assert.IsType(t,
		Args{}, a)
}

func TestArgs_Bool(t *testing.T) {
	for _, v := range [][3]interface{}{
		{&[]bool{false}[0], false, "it's false when false"},
		{&[]bool{true}[0], true, "it's true when true"},
	} {
		a := Args{"k": v[0]}
		assert.IsType(t, a.Bool("k"), v[1], v[2])
	}
}

func TestArgs_String(t *testing.T) {
	for _, v := range [][3]interface{}{
		{&[]string{""}[0], "", "it's false when false"},
	} {
		a := Args{"k": v[0]}
		assert.IsType(t, a.String("k"), v[1], v[2])
	}
}

// func TestArgs_Int(t *testing.T) {
// 	for _, v := range [][3]interface{}{
// 		{&[]int{1}[0], 1, "it's false when false"},
// 	} {
// 		a := Args{"k": v[0]}
// 		assert.IsType(t, a.Int("k"), v[1], v[2])
// 	}
// }

func TestArgs_IsEmpty(t *testing.T) {
	for _, v := range [][3]interface{}{
		{nil, true, "is true on nil"},
		{"string", false, "it false on non-empty string"},
		{"", true, "it true on \"\""},
	} {
		a := Args{"k": v[0]}
		assert.Equal(t, a.IsEmpty("k"), v[1], v[2])
	}
}
