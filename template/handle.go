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
func Open(r, w string, stdout bool) (_readers []Reader, _writer Writer) {
	return readers(r), writer(w, stdout)
}

/**
 */
func cWriter(w Writer) {
	if w != os.Stdout {
		w.Close()
	}
}

/**
 */
func cReader(r []Reader) {
	for _, rr := range r {
		rr.Close()
	}
}

// Close closes all the writers, and readers
// This is an optional method as you can open your
// own in anyway you wish to, and pass it.
func Close(r []Reader, w Writer) {
	cWriter(w)
	cReader(r)
}
