// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package args

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
