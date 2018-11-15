// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTemplate(data string) *Template {
	a := NewFlags().Parse()
	t1, _ := ioutil.TempFile("", "")
	t2, _ := ioutil.TempFile("", "")
	defer t1.Close()
	defer t2.Close()

	a["file"], a["output"] = t1.Name(), t2.Name()
	t := NewTemplate(&a)
	if data != "" {
		t.template.Parse(data)
	}

	return t
}

func TestTemplate__trimStr(t *testing.T) {
	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{"", "    ", "strips a blank string"},
		{"string", "  string", "it strips l whitespace"},
		{"string", "string  ", "it strips r whitespace"},
		{"string", "string\n", "it strips newlines"},
	} {
		actual := tt._trimStr(v[1].(string))
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

func TestTemplate__eExists(t *testing.T) {
	os.Setenv("BLANK", "")
	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{true, "HOME", "it's true if exists"},
		{false, "UNKNOWN", "it's false if it doesn't exist"},
		{true, "BLANK", "it's true if blank"},
	} {
		actual := tt._eExists(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}

func TestTemplate__eStr(t *testing.T) {
	os.Setenv("BLANK", "")
	tt := createTemplate("")
	for _, v := range [][3]interface{}{
		{"HOME", "it's a string if it exists"},
		{"UNKNOWN", "it's a string if it doesn't exist"},
		{"BLANK", "it's a string if it's blank"},
	} {
		actual := tt._eStr(v[0].(string))
		assert.IsType(t, "", actual, v[1])
	}
}

func TestTemplate__eBool(t *testing.T) {
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
		actual, _ := tt._eBool(v[1].(string))
		assert.Equal(t, v[0], actual, v[2])
	}
}
