// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestFuncTrimStr(t *testing.T) {
	tt, _ := template.New("Test").Parse("")
	for _, v := range [][3]interface{}{
		{"", "    ", "strips a blank string"},
		{"string", "  string", "it strips l whitespace"},
		{"string", "string  ", "it strips r whitespace"},
		{"string", "string\n", "it strips newlines"},
	} {
		actual := func_trimStr(tt)(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestFuncTemplateExists(t *testing.T) {
	tt, _ := template.New("Test").Parse("{{ define \"hello\" }}world{{ end }}")
	for _, v := range [][3]interface{}{
		{true, "Test", "it's true for itself"},
		{true, "hello", "it's true even for define"},
		{false, "", "it's false if blank"},
	} {
		actual := func_templateExists(tt)(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestFuncEExists(t *testing.T) {
	os.Setenv("BLANK", "")
	tt, _ := template.New("Test").Parse("")
	for _, v := range [][3]interface{}{
		{true, "HOME", "it's true if exists"},
		{false, "UNKNOWN", "it's false if it doesn't exist"},
		{true, "BLANK", "it's true if blank"},
	} {
		actual := func_eExists(tt)(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestFuncEStr(t *testing.T) {
	os.Setenv("BLANK", "")
	tt, _ := template.New("Test").Parse("")
	for _, v := range [][3]interface{}{
		{"HOME", "it's a string if it exists"},
		{"UNKNOWN", "it's a string if it doesn't exist"},
		{"BLANK", "it's a string if it's blank"},
	} {
		actual := func_eStr(tt)(v[0].(string))
		assert.IsType(t, "", actual, v[1])
	}
}

func TestFuncEBool(t *testing.T) {
	os.Setenv("TRUE_1", "1")
	os.Setenv("TRUE_true", "true")
	os.Setenv("FALSE_false", "false")
	os.Setenv("FALSE_0", "0")
	os.Setenv("BLANK", "")

	tt, _ := template.New("Test").Parse("")
	for _, v := range [][3]interface{}{
		{false, "FALSE_0", "it's false if it's 0"},
		{true, "TRUE_true", "it's true if it's true"},
		{false, "UNKNOWN", "it's false if it doesn't exit"},
		{false, "HOME", "it's false if it exists and isn't true/1"},
		{false, "FALSE_false", "it's false if it's false"},
		{false, "BLANK", "it's false if it's blank"},
		{true, "TRUE_1", "it's true if it's 1"},
	} {
		actual, _ := func_eBool(tt)(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}
