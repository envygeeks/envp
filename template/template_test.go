// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func (t *TestReader) Name() string { return t._name }
func (t *TestWriter) Name() string { return t._name }
func (t *TestReader) Close() error { return nil }
func (t *TestWriter) Close() error { return nil }

/**
 */
type TestWriter struct {
	_name string
	*strings.Builder
}

/**
 */
type TestReader struct {
	_name string
	*strings.Reader
}

/**
 */
func TestParseFile(t *testing.T) {
	type TestStruct struct {
		expected    string
		description string
		name        string
	}

	for _, test := range []TestStruct{
		TestStruct{
			expected:    "Hello World",
			description: "it's not nil",
			name:        "hello",
		},
	} {
		var str strings.Builder

		template := New(false)
		reader := &TestReader{
			Reader: strings.NewReader(test.expected),
			_name:  test.name,
		}

		template.ParseFile(reader)
		atemplate := template.Lookup(test.name)
		if assert.NotNil(t, atemplate) {
			atemplate.Execute(&str, "")

			actual := str.String()
			assert.Equal(t, test.expected, actual,
				test.description)
		}
	}
}

func TestExec(t *testing.T) {
	type TestStruct struct {
		expected    string
		description string
		name        string
		use         string
	}

	for _, test := range []TestStruct{
		TestStruct{
			expected:    "Hello World: 1",
			description: "it returns a byte",
			name:        "hello-1",
		},
		TestStruct{
			name:        "hello-2",
			expected:    "Hello World: 2",
			description: "it works with use",
			use:         "hello-2",
		},
	} {
		template := New(false)
		reader := &TestReader{
			Reader: strings.NewReader(test.expected),
			_name:  test.name,
		}

		if test.use != "" {
			template.Use(reader)
			expected, actual := test.name, template.use
			assert.Equal(t, expected, actual,
				test.description)
		}

		template.ParseFile(reader)
		actual := string(template.Compile())
		assert.Equal(t, test.expected,
			actual)
	}
}

func TestWrite(t *testing.T) {
	type TestStruct struct {
		expected    string
		description string
		name        string
	}

	for _, test := range []TestStruct{
		TestStruct{
			expected:    "hello",
			description: "it writes the template",
			name:        "test1",
		},
	} {
		template := New(false)
		reader := &TestReader{
			Reader: strings.NewReader(test.expected),
			_name:  test.name,
		}

		template.ParseFile(reader)
		writer := &TestWriter{
			Builder: new(strings.Builder),
		}

		out := template.Compile()
		template.Write(out, writer)
		actual := writer.String()
		assert.Equal(t, test.expected, actual,
			test.description)
	}
}

func TestOpen(t *testing.T) {
	fs := afero.NewOsFs()
	readf, _ := afero.TempFile(fs, "", "test-open-returns-stdout")
	writf, _ := afero.TempFile(fs, "", "test-open-returns-stdout")
	defer func() { readf.Close(); fs.Remove(readf.Name()) }()
	defer func() { writf.Close(); fs.Remove(writf.Name()) }()
	reader, writer := Open(readf.Name(), writf.Name())
	assert.NotEmpty(t, reader)
	assert.NotNil(t, writer)
}

/**
 */
type TestCloseW struct {
	io.Writer
	CRan bool
}

/**
 */
type TestCloseR struct {
	CRan bool
	io.Reader
}

func (r *TestCloseR) Name() (s string) { return }
func (w *TestCloseW) Name() (s string) { return }
func (r *TestCloseR) Close() (e error) { r.CRan = true; return }
func (w *TestCloseW) Close() (e error) {
	w.CRan = true
	return
}

func TestClose(t *testing.T) {
	fs := new(afero.MemMapFs)
	readf, _ := afero.TempFile(fs, "", "test-close")
	writf, _ := afero.TempFile(fs, "", "test-close")
	defer func() { readf.Close(); fs.Remove(readf.Name()) }()
	defer func() { writf.Close(); fs.Remove(writf.Name()) }()
	writer := &TestCloseW{Writer: readf}
	reader := &TestCloseR{Reader: writf}
	Close([]Reader{reader}, writer)

	assert.True(t, writer.CRan)
	assert.True(t, reader.CRan)
}
