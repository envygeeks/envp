// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
)

type BoolArg []interface{}
type Args map[string]interface{}
type StringArg []interface{}
type Flags []interface{}

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
	o, ok := a[k].(string)
	if ok {
		return o
	}

	v := *a[k].(*string)
	return string(v)
}

// IsEmpty Allows you to check if a value was
// given, this means that it wasn't nil, that it
// exists on the map, and that it's not empty.
func (a Args) IsEmpty(k string) bool {
	if v, ok := a[k]; !ok || v == nil || v == "" {
		return true
	}

	return false
}

// NewFlags provides the default flags, and
// allows somebody who wishes to use as a base to
// add their own flags, by just append()
func NewFlags() *Flags {
	return &Flags{
		StringArg{"file", "", "the file, or dir"},
		BoolArg{"glob", false, "search, and use a dir full of *.gohtml"},
		StringArg{"output", "", "the file to write to"},
		BoolArg{"stdout", false, "print to stdout"},
		BoolArg{"debug", false, "debug output"},
	}
}

// Parse parses our flags, and returns them,
// we just encapsulate this to make life a little
// easier at the end of the day, when working.
func (f Flags) Parse() (a Args) {
	a = make(Args)
	for _, v := range f {
		switch v.(type) {
		case BoolArg:
			vv := v.(BoolArg)
			k, def, dsc := vv[0].(string), vv[1].(bool), vv[2].(string)
			if flag.Lookup(k) == nil {
				a[k] = flag.Bool(k, def, dsc)
				continue
			}

			fv := flag.Lookup(k).Value
			ov := fv.(flag.Getter).Get()
			a[k] = ov.(bool)

		case StringArg:
			vv := v.(StringArg)
			k, def, dsc := vv[0].(string), vv[1].(string), vv[2].(string)
			if flag.Lookup(k) == nil {
				a[k] = flag.String(k, def, dsc)
				continue
			}

			fv := flag.Lookup(k).Value
			ov := fv.(flag.Getter).Get()
			a[k] = ov.(string)
		}
	}

	if !flag.Parsed() {
		flag.Parse()
	}

	return
}
