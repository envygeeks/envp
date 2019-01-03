package cmd

import (
	"regexp"
	"strings"

	upstream "github.com/envygeeks/envp/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type rootCmd struct {
	*cobra.Command
}

var (
	sR = `(?m)^\t+`
	sL = regexp.MustCompile(sR).
		ReplaceAllString
)

func tS(s string) string {
	return strings.TrimSpace(sL(s, ""))
}

var (
	Root = (&rootCmd{
		&cobra.Command{
			Short: "Build your configs",
			Use:   "envp",

			Long: tS(`
				Build your configuration files with helpers, and access
				to the current env, so that you can shim configuration files
				in a Docker image when they do no support such mechanisms.
			`),
		},
	}).Init()
)

// disableHelp disables `cmd help`
func (r *rootCmd) disableHelp() {
	r.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
}

// Init sets everything up, it disables
// the help command, it also sets the flags,
// make sure start is wrapped in, and more
func (r *rootCmd) Init() *rootCmd {
	r.disableHelp()

	r.PreRun = r.PreStart
	r.Flags().String("write-to", "", "write to (stdout)")
	r.PersistentFlags().Bool("debug", false, "verbose debug output")
	r.Flags().StringArray("file", []string{}, "files to read in as templates")
	r.Flags().Bool("version", false, "the current app version")
	r.Run = r.Start
	return r
}

// files pulls the files, and checks
func (r *rootCmd) files() []string {
	files, err := r.Flags().GetStringArray("file")
	if err != nil {
		logrus.Fatalln(err)
	}

	return files
}

// writeTo pulls down write-to
func (r *rootCmd) writeTo() string {
	writeTo, err := r.Flags().GetString("write-to")
	if err != nil {
		logrus.Fatalln(err)
	}

	return writeTo
}

// preStart runs stuff before start
func (r *rootCmd) PreStart(*cobra.Command, []string) {
	logrus.SetLevel(logrus.WarnLevel)
	if t, err := r.Flags().GetBool("debug"); err == nil && t {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// Start and run the command
func (r *rootCmd) Start(*cobra.Command, []string) {
	ver, err := r.Flags().GetBool("version")
	if err != nil {
		logrus.Fatalln(err)
	} else {
		if ver {
			Info.PrintSimple()
		}
	}

	template := upstream.New()
	writeTo, files := r.writeTo(), r.files()
	readers, writer := upstream.Open(files, writeTo)
	defer upstream.Close(readers, writer)
	template.ParseFiles(readers)
	if len(readers) == 1 {
		template.Use(readers[0])
	}

	byte := template.Compile()
	template.Write(byte, writer)
}
