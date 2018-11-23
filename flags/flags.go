// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package flags

import (
	"flag"

	"github.com/envygeeks/envp/flags/args"
)

// List is a list of args
type List []interface {
	/**
	 * A list of:
	 *   - []String
	 *   - []Bool
	 */
}

// String is a string type argument
// exp: -string=val will be a string
// when it reaches you.
type String struct {
	Name, Default, Description string
}

// Bool is a bool type argument
// exp: -bool=true|1, or -bool=false|0
// and will be a bool when it
// finally reaches you
type Bool struct {
	Name, Description string
	Default           bool
}

var (
	ran      = false
	parsed   *args.Args
	defaults = &List{
		String{
			Name:        "file",
			Description: "the file, or dir",
			Default:     "",
		},
		String{
			Name:        "output",
			Description: "the file to write to",
			Default:     "",
		},
		Bool{
			Default:     false,
			Description: "debug output",
			Name:        "debug",
		},
	}
)

// Parse parses our flags
func Parse() *args.Args {
	if !ran || parsed == nil {
		ran = true
		parsed = defaults.Parse()
		return parsed
	}

	return parsed
}

// Parse parses our flags, and returns them,
// we just encapsulate this to make life a little
// easier at the end of the day, when working.
func (f List) Parse() *args.Args {
	a := args.New()
	for _, v := range f {
		switch v.(type) {
		case Bool:
			vv := v.(Bool)
			if flag.Lookup(vv.Name) == nil {
				a[vv.Name] = flag.Bool(vv.Name, vv.Default,
					vv.Description)

				continue
			}

			fv := flag.Lookup(vv.Name).Value
			ov := fv.(flag.Getter).Get()
			a[vv.Name] = ov.(bool)

		case String:
			vv := v.(String)
			if flag.Lookup(vv.Name) == nil {
				a[vv.Name] = flag.String(vv.Name, vv.Default,
					vv.Description)

				continue
			}

			fv := flag.Lookup(vv.Name).Value
			ov := fv.(flag.Getter).Get()
			a[vv.Name] = ov.(string)
		}
	}

	if !flag.Parsed() {
		flag.Parse()
	}

	return &a
}
