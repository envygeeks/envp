// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package args

// Args is a map
type Args map[string]interface {
	/**
	 * Your args here
	 * On New
	 */
}

// New creates a new instance of Args for
// flags to be placed on as they are parsed
// this is passed around to template.
func New() Args {
	return make(Args)
}
