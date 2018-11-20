// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package template

import (
	"os"
)

// Open opens all the readers, and writers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Open(r, w string, stdout bool) (_readers []*os.File, _writer *os.File) {
	return readers(r), writer(w, stdout)
}

// Close closes all the writers, and readers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Close(r []*os.File, w *os.File) {
	w.Close()
	for _, rr := range r {
		rr.Close()
	}
}
