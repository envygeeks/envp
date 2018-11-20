// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ts struct {
	a interface{}
	e interface{}
	d string
}

type testReader struct {
	_name string
	*strings.Reader
}

func (t *testReader) Name() string {
	return t._name
}

func TestParse(t *testing.T) {
	tt := New(false)
	rr := &testReader{
		Reader: strings.NewReader("hello"),
		_name:  "test1",
	}

	tt.Parse(rr)
	ttt := tt.template.Lookup("test1")
	assert.NotNil(t, ttt)
}

func TestExec(t *testing.T) {
	tt := New(false)
	rr := &testReader{
		Reader: strings.NewReader("hello"),
		_name:  "test1",
	}

	tt.Parse(rr)
	a := string(tt.Exec())
	e := "hello"

	assert.Equal(t, e, a)
}

func TestWrite(t *testing.T) {
	tt := New(false)
	rr := &testReader{
		Reader: strings.NewReader("hello"),
		_name:  "test1",
	}

	tt.Parse(rr)
	var s strings.Builder
	tt.Write(tt.Exec(), &s)
	a := s.String()
	e := "hello"

	assert.Equal(t, e, a)
}
