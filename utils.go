// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// expand expands a path just throwing the error
// upstream as a fatal error if we cannot expand it
// instead of passing the error downstream.
func expand(s string) string {
	f, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

// isExist lets you know if a path exists, it
// doesn't really matter if it's a directory or
// a file, just that it exist.
func isExist(s string) bool {
	if _, err := os.Stat(s); !os.IsNotExist(err) {
		return true
	}

	return false
}
