// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package logger

import (
	"io"
	"log"
	"os"
)

var (
	normal = log.New(os.Stderr, "", 0)
	always = log.New(os.Stderr, "", 0)
)

/**
 */
func SetOutput(w io.Writer)    { normal.SetOutput(w) }
func Println(v ...interface{}) { normal.Println(v...) }
func Fatalln(v ...interface{}) { always.Fatalln(v...) }
func Panicln(v ...interface{}) {
	always.Panicln(v...)
}

/**
 */
func Printf(format string, v ...interface{}) { normal.Printf(format, v...) }
func Fatalf(format string, v ...interface{}) { always.Fatalf(format, v...) }
func Panicf(format string, v ...interface{}) {
	always.Panicf(format, v...)
}

/**
 */
func Print(v ...interface{}) { normal.Print(v...) }
func Fatal(v ...interface{}) { always.Fatal(v...) }
func Panic(v ...interface{}) {
	always.Panic(v...)
}
