// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package args

// IsEmpty Allows you to check if a value was
// given, this means that it wasn't nil, that it
// exists on the map, and that it's not empty.
func (a Args) IsEmpty(k string) bool {
	if v, ok := a[k]; !ok || v == nil || v == "" {
		return true
	}

	return false
}
