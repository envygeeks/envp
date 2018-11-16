// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate__reindent(t *testing.T) {
	s, e := "\n\n\t1\n\t  2\n\t3", "1\n  2\n3"
	tt := createTemplate("")
	actual := tt._reindent(s)
	assert.Equal(t, e, actual,
		"it reindents")
}

func TestTemplate__trimEmpty(t *testing.T) {
	s, e := "\n\n\n\n1\r\n\r\n", "1"
	tt := createTemplate("")
	actual := tt._trimEdges(s)
	assert.Equal(t, e, actual,
		"it strips")
}

func TestTemplate__space(t *testing.T) {
	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{" 1", "1", "it works for simple strings"},
		{" 1", "   1 ", "it works for strings with space"},
		{" 1", "1\n", "it works with newline"},
	} {
		actual := tt._space(v[1].(string), 1)
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestTemplate__templateExists(t *testing.T) {
	tt := createTemplate("{{ define \"hello\" }}world{{ end }}")
	for _, v := range [][3]interface{}{
		{true, "envp", "it's true for itself"},
		{true, "hello", "it's true even for define"},
		{false, "", "it's false if blank"},
	} {
		actual := tt._templateExists(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestTemplate__envExists(t *testing.T) {
	os.Setenv("BLANK", "")
	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{true, "HOME", "it's true if exists"},
		{false, "UNKNOWN", "it's false if it doesn't exist"},
		{true, "BLANK", "it's true if blank"},
	} {
		actual := tt._envExists(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestTemplate__env(t *testing.T) {
	os.Setenv("BLANK", "")
	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{"HOME", "it's a string if it exists"},
		{"UNKNOWN", "it's a string if it doesn't exist"},
		{"BLANK", "it's a string if it's blank"},
	} {
		actual := tt._env(v[0].(string))
		assert.IsType(t, "", actual, v[1])
	}
}

func TestTemplate__boolEnv(t *testing.T) {
	os.Setenv("TRUE_1", "1")
	os.Setenv("TRUE_true", "true")
	os.Setenv("FALSE_false", "false")
	os.Setenv("FALSE_0", "0")
	os.Setenv("BLANK", "")

	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{false, "FALSE_0", "it's false if it's 0"},
		{true, "TRUE_true", "it's true if it's true"},
		{false, "UNKNOWN", "it's false if it doesn't exit"},
		{false, "HOME", "it's false if it exists and isn't true/1"},
		{false, "FALSE_false", "it's false if it's false"},
		{false, "BLANK", "it's false if it's blank"},
		{true, "TRUE_1", "it's true if it's 1"},
	} {
		a := tt._boolEnv(v[1].(string))
		assert.Equal(t, v[0], a, v[2])
	}
}
