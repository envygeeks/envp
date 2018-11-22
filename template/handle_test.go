// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	fs := afero.NewOsFs()
	r, _ := afero.TempFile(fs, "", "test-open-returns-stdout")
	w, _ := afero.TempFile(fs, "", "test-open-returns-stdout")
	defer func() { r.Close(); fs.Remove(r.Name()) }()
	defer func() { w.Close(); fs.Remove(w.Name()) }()
	rs, ws := Open(r.Name(), w.Name(), true)
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
