// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
)

type BoolSlice []interface{}
type Args map[string]interface{}
type StringSlice []interface{}
type IntSlice []interface{}
type ArgSlice []interface{}

var (
	_cliFlags = ArgSlice{
		StringSlice{"file", "", "the file, or dir"},
		BoolSlice{"glob", false, "search, and use a dir full of *.gohtml"},
		StringSlice{"output", "", "the file to write to"},
		BoolSlice{"stdout", false, "print to stdout"},
	}
)

func (a Args) Bool(k string) bool     { return bool(*a[k].(*bool)) }
func (a Args) String(k string) string { return string(*a[k].(*string)) }
func (a Args) Int(k string) int       { return int(*a[k].(*int)) }

// IsEmpty Allows you to check if a value was
// given, this means that it wasn't nil, that it
// exists on the map, and that it's not empty.
func (a Args) IsEmpty(k string) bool {
	if v, ok := a[k]; !ok || v == nil || v == "" {
		return true
	}

	return false
}

// flags parses our flags, and returns them
// we just encapsulate this to make life a little
// easier at the end of the day, when working.
func flags() (a Args) {
	a = make(map[string]interface{})
	for _, v := range _cliFlags {
		switch v.(type) {
		case BoolSlice:
			vv := v.(BoolSlice)
			a[vv[0].(string)] = flag.Bool(vv[0].(string),
				vv[1].(bool), vv[2].(string))

		case StringSlice:
			vv := v.(StringSlice)
			a[vv[0].(string)] = flag.String(vv[0].(string),
				vv[1].(string), vv[2].(string))

		case IntSlice:
			vv := v.(IntSlice)
			a[vv[0].(string)] = flag.Int(vv[0].(string),
				vv[1].(int), vv[2].(string))
		}
	}

	flag.Parse()
	return
}
