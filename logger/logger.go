// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	normal = log.New(os.Stderr, "", 0)
	always = log.New(os.Stderr, "", 0)
)

// Disable the logger in testing
// The output just gets in the way
// we don't really need it.
func init() {
	if strings.HasSuffix(os.Args[0], ".test") {
		normal.SetOutput(ioutil.Discard)
	}
}

// SetOutput allows you to set the output
// This is ignored for "always" which is tapped
// on things like Fatal, and Panic.
func SetOutput(w io.Writer) {
	normal.SetOutput(w)
}

// Println forwards
// This can be turned off.
// -> normal.Println
func Println(v ...interface{}) {
	normal.Println(v...)
}

// Fatalln forwards
// This cannot be turned off.
// -> always.Fatalln
func Fatalln(v ...interface{}) {
	always.Fatalln(v...)
}

// Panicln forwards
// This cannot be turned off.
// -> always.Panicln
func Panicln(v ...interface{}) {
	always.Panicln(v...)
}

// Printf forwards
// This can be turned off.
// -> normal.Printf
func Printf(format string, v ...interface{}) {
	normal.Printf(format,
		v...)
}

// Fatalf forwards
// This cannot be turned off.
// -> always.Fatalf
func Fatalf(format string, v ...interface{}) {
	always.Fatalf(format,
		v...)
}

// Panicf forwards
// This cannot be turned off.
// -> always.Panicf
func Panicf(format string, v ...interface{}) {
	always.Panicf(format,
		v...)
}

// Panic forwards
// This cannot be turned off.
// -> always.Panic
func Panic(v ...interface{}) {
	always.Panic(v...)
}
