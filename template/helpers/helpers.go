// Copyright 2018 Jordon Bedwell. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package helpers

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/envygeeks/envp/logger"
)

// Helpers holds the template
// We use it for some template funcs
// Some don't use it at all.
type Helpers struct {
	template *template.Template
}

// EnvExists allows you to check if a var exists
func (h *Helpers) EnvExists(s string) bool {
	s = strings.ToUpper(s)
	if _, ok := os.LookupEnv(s); ok {
		return true
	}

	return false
}

// Env allows you to pull out a string var
func (h *Helpers) Env(s string) string {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		return v
	}

	return ""
}

// BoolEnv allows you to pull out a var as bool
func (h *Helpers) BoolEnv(s string) bool {
	s = strings.ToUpper(s)
	if v, ok := os.LookupEnv(s); ok {
		obool, err := strconv.ParseBool(v)
		if err != nil {
			logger.Println(err)
			return false
		}

		return obool
	}

	return false
}

// AddSpace adds a space to the beginning of a string
func (h *Helpers) AddSpace(s string, n int) string {
	s = strings.TrimSpace(s)
	space := strings.Repeat(" ", n)
	return space + s
}

const (
	indentRegex     = `(?m)^[ \t]{%d}`
	stripEdgesRegex = `(?m)\A[ \t]*$[\r\n]*|[\r\n]+[ \t]*\z`
	stripEmptyRegex = `(?m)^[ \t]+$`
	preindentRegex  = `(?m)^[ \t]+`
)

func indent(h *Helpers, s string, size int) string {
	re, s, indent := regexp.MustCompile(preindentRegex), h.Strip(s), -1
	for _, v := range re.FindAllString(s, -1) {
		if l := len(v); indent == -1 || l < indent {
			indent = l
			continue
		}
	}

	if indent > -1 {
		repeated := fmt.Sprintf(indentRegex, indent)
		re := regexp.MustCompile(repeated)
		if size > -1 {
			i := strings.Repeat(" ", size)
			replaced := re.ReplaceAllString(s, i)
			return replaced
		}

		replaced := re.ReplaceAllString(s, "")
		return replaced
	}

	return s
}

// FixIndentation strips indentation to the edge
func (h *Helpers) FixIndentation(s string) string {
	out := indent(h, s, -1)
	return out
}

// Indent strips each lines indentation to the
// edge, and then indents each line to the size you set
func (h *Helpers) Indent(s string, size uint) string {
	out := indent(h, s, int(size))
	return out
}

// Strip removes empty lines around a string
func (h *Helpers) Strip(s string) string {
	re1 := regexp.MustCompile(stripEmptyRegex)
	re2 := regexp.MustCompile(stripEdgesRegex)
	s = re1.ReplaceAllString(s, "")
	s = re2.ReplaceAllString(s, "")
	return s
}

// TemplateString pulls a template as a string
func (h *Helpers) TemplateString(s string) string {
	if template := h.template.Lookup(s); template != nil {
		var str strings.Builder

		if err := template.Execute(&str, h.template); err != nil {
			logger.Fatalln(err)
		}

		out := str.String()
		return out
	}

	// Bad template given.
	logger.Fatalf("Unable to find %s", s)
	return ""
}

// IndentedTemplate indents a template.
func (h *Helpers) IndentedTemplate(s string, size uint) string {
	s = h.TemplateString(s)
	s = h.Indent(s, size)
	return s
}

// TemplateWithNewLine returns a template with
// a newline if the template returned is not empty
func (h *Helpers) TemplateWithNewLine(s string) string {
	s = h.FixIndentedTemplate(s)
	if s != "" {
		return "\n" + s
	}

	return s
}

// IndentedTemplateWithNewLine adds a newline, and indents
func (h *Helpers) IndentedTemplateWithNewLine(s string, size uint) string {
	s = h.IndentedTemplate(s, size)
	if s != "" {
		return "\n" + s
	}

	return s
}

// StrippedTemplate trims empty lines, and edges.
func (h *Helpers) StrippedTemplate(s string) string {
	s = h.TemplateString(s)
	s = h.Strip(s)
	return s
}

// FixIndentedTemplate strips the indentation to the edge
func (h *Helpers) FixIndentedTemplate(s string) string {
	s = h.TemplateString(s)
	s = h.FixIndentation(s)
	return s
}

// TemplateExists checks if a template exists
func (h *Helpers) TemplateExists(s string) bool {
	logger.Printf("looking for template %s", s)
	if template := h.template.Lookup(s); template != nil {
		return true
	}

	return false
}

var (
	letters    = []rune("01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersLen = int64(len(letters))
)

func rngCheck() {
	buf := make([]byte, 1)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		log.Fatalln(err)
	}
}

// RandomPassword generates a password with
// cryptographically derived random numbers
func (h *Helpers) RandomPassword(size uint) string {
	rngCheck()

	rune := make([]rune, size)
	for i := range rune {
		n, err := rand.Int(rand.Reader, big.NewInt(lettersLen))
		if err != nil {
			log.Fatalln(err)
		}

		idx := n.Int64()
		letter := letters[idx]
		rune[i] = letter
	}

	out := string(rune)
	return out
}

// New creates a new Funcs, and registers them
func New(t *template.Template) *Helpers {
	helpers := (&Helpers{template: t}).Register()
	return helpers
}

// Register registers the funcs
func (h *Helpers) Register() *Helpers {
	logger.Println("registering all the helpers")
	h.template.Funcs(template.FuncMap{
		"split":                       strings.Split,
		"chomp":                       strings.Trim,
		"indent":                      h.Indent,
		"addSpace":                    h.AddSpace,
		"random_password":             h.RandomPassword,
		"templateString":              h.TemplateString,
		"strippedTemplate":            h.StrippedTemplate,
		"fixIndentedTemplate":         h.FixIndentedTemplate,
		"indentedTemplateWithNewline": h.IndentedTemplateWithNewLine,
		"templateWithNewline":         h.TemplateWithNewLine,
		"indentedTemplate":            h.IndentedTemplate,
		"templateExists":              h.TemplateExists,
		"fixIndentation":              h.FixIndentation,
		"envExists":                   h.EnvExists,
		"bool_env":                    h.BoolEnv,
		"strip":                       h.Strip,
		"env":                         h.Env,
	})

	return h
}
