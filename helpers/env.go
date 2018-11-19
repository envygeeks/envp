// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// EnvExists allows you to check if a var exists
// in your current environment, we do not alter it so
// make sure you use FULLCAP if necessary.
func (h *Helpers) EnvExists(s string) bool {
	s = strings.ToUpper(s)
	if _, ok := os.LookupEnv(s); ok {
		return true
	}

	return false
}

// Env allows you to pull out a string env var
func (h *Helpers) Env(s string) string {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		return v
	}

	return ""
}

// BoolEnv allows you to pull out a env var as a
// bool, following the same rules as strconv.ParseBool
// where 1, true are true, and all else is false
func (h *Helpers) BoolEnv(s string) bool {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		vv, err := strconv.ParseBool(v)
		if err != nil {
			log.Warn(err)
			return false
		}

		return vv
	}

	return false
}
