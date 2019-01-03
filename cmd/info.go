package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type infoCmd struct {
	*cobra.Command
}

var (
	date time.Time

	commit  string
	strDate string
	version string
	url     string

	Info = (&infoCmd{
		&cobra.Command{
			Short: "Show build info",
			Use:   "info",
		},
	}).Init()
)

func init() {
	Root.AddCommand(Info.Command)
}

func (infoCmd) setCommit() {
	if commit == "" {
		commit = "HEAD"
		out, err := exec.Command("git", "rev-parse", "--verify", "HEAD").Output()
		if err != nil {
			logrus.Fatalln(err)
		}

		commit += " @ "
		str := strings.TrimSpace(string(out))
		commit += str
	}
}

func (infoCmd) setVersion() {
	if version == "" {
		out, err := exec.Command("git", "describe", "--abbrev=0").Output()
		if err != nil {
			logrus.Fatalln(err)
		}

		str := string(out)
		str = strings.TrimSpace(str)
		version = str
	}
}

// gitUrl will convert a git@github
// to a proper url that can be linked in
// the terminal for your use
func gitUrl(s string) string {
	s = strings.Replace(s, ":", "/", -1)
	re := regexp.MustCompile(`(^origin[\s\t]+git@|\.git \((fetch|push)\)$)`)
	byte := re.ReplaceAll([]byte(s), []byte(""))
	s = "https://" + string(byte)
	return s
}

func (infoCmd) setUrl() {
	if url == "" {
		out, err := exec.Command("git", "remote", "-v").Output()
		if err != nil {
			logrus.Fatalln(err)
		}

		reader := bufio.NewReader(bytes.NewReader(out))
		for {
			line, err := reader.ReadString(0x0a)
			if err != nil {
				if err == io.EOF {
					break
				}

				logrus.Fatalln(err)
			}

			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "origin") {
				url = gitUrl(line)
				break
			}
		}
	}

	if url == "" {
		// Are we even a Git repository?
		logrus.Fatalln("unable to set url")
	}
}

func (infoCmd) setDate() {
	if strDate != "" {
		intDate, err := strconv.ParseInt(strDate, 10, 64)
		if err != nil {
			logrus.Fatalln(err)
		}

		date = time.Unix(intDate, 0)
		return
	}

	date = time.Now()
	return
}

// Init finishes initializing verCmd
func (i *infoCmd) Init() *infoCmd {
	i.setDate()
	i.setCommit()
	i.setVersion()
	i.setUrl()

	i.Run = i.Start
	i.Flags().Bool("simple", false, "only print the version")
	return i
}

// PrintComplex prints more verbose
// things like commit, build date, and
// other useful information for you
func (infoCmd) PrintComplex() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Url: %s\n", url)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Date: %d-%02d-%02d\n",
		date.Year(), date.Month(),
		date.Day())
}

// PrintSimple only prints the version
// this format will be v0.0.0.alpha
func (infoCmd) PrintSimple() {
	if version != "" {
		fmt.Println(version)
	}
}

// Start runs the command
func (i *infoCmd) Start(_ *cobra.Command, args []string) {
	simple, err := i.Flags().GetBool("simple")
	if err != nil {
		logrus.Fatalln(err)
	} else {
		if simple {
			i.PrintSimple()
			return
		}
	}

	i.PrintComplex()
}
