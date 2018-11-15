// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	log "github.com/sirupsen/logrus"
)

/**
 */
func main() {
	a := NewFlags().Parse()
	l := log.WarnLevel
	if a.Bool("debug") {
		l = log.DebugLevel
	}

	log.SetLevel(l)
	NewTemplate(&a).Run()
}
