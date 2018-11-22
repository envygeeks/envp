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

// Bool pulls a value as a boolean, this
// supports *bool, and bool, because either
// can be given depending on flags
func (a Args) Bool(k string) bool {
	o, ok := a[k].(bool)
	if ok {
		return o
	}

	v := *a[k].(*bool)
	return bool(v)
}

// String pulls a value as a boolean, this
// supports *bool, and bool, because either
// can be given depending on flags
func (a Args) String(k string) string {
	v, ok := a[k]
	if !ok {
		return ""
	}

	if s, ok := v.(string); !ok {
		if s, ok := v.(*string); ok && s != nil {
			return string(*s)
		}
	} else {
		return s
	}

	return ""
}

// IsEmpty Allows you to check if a value was
// given, this means that it wasn't nil, that it
// exists on the map, and that it's not empty.
func (a Args) IsEmpty(k string) bool {
	if v := a.String(k); v == "" {
		return true
	}

	return false
}
