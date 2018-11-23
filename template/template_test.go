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
func TestParse(t *testing.T) {
	type TestStruct struct {
		expected    string
		description string
		name        string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    "Hello World",
			description: "it's not nil",
			name:        "hello",
		},
	} {
		tt := New(false)
		tr := &TestReader{
			Reader: strings.NewReader(ts.expected),
			_name:  ts.name,
		}

		tt.Parse(tr)
		var s strings.Builder
		ttt := tt.template.Lookup(ts.name)
		if assert.NotNil(t, ttt) {
			ttt.Execute(&s, "")
			actual := s.String()
			assert.Equal(t, ts.expected, actual,
				ts.description)
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

	for _, ts := range []TestStruct{
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
		tt := New(false)
		tr := &TestReader{
			Reader: strings.NewReader(ts.expected),
			_name:  ts.name,
		}

		if ts.use != "" {
			tt.Use(tr)
			expected, actual := ts.name, tt.use
			assert.Equal(t, expected, actual,
				ts.description)
		}

		tt.Parse(tr)
		actual := string(tt.Exec())
		assert.Equal(t, ts.expected,
			actual)
	}
}

func TestWrite(t *testing.T) {
	type TestStruct struct {
		expected    string
		description string
		name        string
	}

	for _, ts := range []TestStruct{
		TestStruct{
			expected:    "hello",
			description: "it writes the template",
			name:        "test1",
		},
	} {
		tt := New(false)
		tr := &TestReader{
			Reader: strings.NewReader(ts.expected),
			_name:  ts.name,
		}

		tt.Parse(tr)
		tw := &TestWriter{
			Builder: new(strings.Builder),
		}

		o := tt.Exec()
		tt.Write(o, tw)
		actual := tw.String()
		assert.Equal(t, ts.expected, actual,
			ts.description)
	}
}

func TestOpen(t *testing.T) {
	fs := afero.NewOsFs()
	r, _ := afero.TempFile(fs, "", "test-open-returns-stdout")
	w, _ := afero.TempFile(fs, "", "test-open-returns-stdout")
	defer func() { r.Close(); fs.Remove(r.Name()) }()
	defer func() { w.Close(); fs.Remove(w.Name()) }()
	rs, ws := Open(r.Name(), w.Name())
	assert.NotEmpty(t, rs)
	assert.NotNil(t, ws)
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
	r, _ := afero.TempFile(fs, "", "test-close")
	w, _ := afero.TempFile(fs, "", "test-close")
	defer func() { r.Close(); fs.Remove(r.Name()) }()
	defer func() { w.Close(); fs.Remove(w.Name()) }()
	wt := &TestCloseW{Writer: r}
	rt := &TestCloseR{Reader: w}
	Close([]Reader{rt}, wt)

	assert.True(t, wt.CRan)
	assert.True(t, rt.CRan)
}
