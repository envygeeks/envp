// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package main

import (
	"flag"

	"github.com/envygeeks/envp/args"
)

type BoolArg []interface{}
type StringArg []interface{}
type Flags []interface{}

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
func (f Flags) Parse() (a args.Args) {
	a = args.New()
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
