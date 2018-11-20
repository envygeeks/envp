// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package utils

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// IsDir tells you if a path is a directory
func IsDir(s string) bool {
	finfo, err := os.Stat(s)
	if err != nil {
		log.Warningln(err)
		return false
	}

	return finfo.IsDir()
}

// Expand expands a path just throwing the error
// upstream as a fatal error if we cannot expand it
// instead of passing the error downstream.
func Expand(s string) string {
	f, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

// IsExist lets you know if a path exists, it
// doesn't really matter if it's a directory or
// a file, just that it exist.
func IsExist(s string) bool {
	if _, err := os.Stat(s); !os.IsNotExist(err) {
		return true
	}

	return false
}
