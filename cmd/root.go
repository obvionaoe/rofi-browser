/*
Copyright © 2024 Luís Guimarães

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"cmp"
	"github.com/spf13/cobra"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	browser      string
	sortProfiles bool
	rofiCmd      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rofi-browser",
	Short: "A simple profile picker for Mozilla Firefox based web browsers. Uses rofi.",
	Long: `rofi-browser is a rofi wrapper that provides a simple way to launch
specific profiles in Mozilla Firefox and some of its forks.`,
	Run: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(
		&browser,
		"browser",
		"b",
		"firefox",
		"The browser to use",
	)
	rootCmd.Flags().BoolVarP(
		&sortProfiles,
		"sort",
		"s",
		false,
		"Sort the profiles alphabetically",
	)
	rootCmd.Flags().StringVarP(
		&rofiCmd,
		"rofi",
		"r",
		"rofi -dmenu -p \"Profile\" -no-custom",
		"Command to launch rofi, including flags",
	)
}

func run(cmd *cobra.Command, args []string) {
	profilesFile := ""
	if browser == "firefox" {
		profilesFile = "~/.mozilla/firefox/profiles.ini"
	} else {
		profilesFile = "~/.librewolf/profiles.ini"
	}

	profilesFile, err := expandTilde(profilesFile)
	if err != nil {
		log.Fatal(err)
	}

	profiles, err := getProfiles(profilesFile)
	if err != nil {
		log.Fatal(err)
	}

	if sortProfiles {
		slices.SortFunc(profiles, func(a, b string) int {
			return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
		})
	}

	selectedProfile, err := runRofi(profiles)
	if err != nil {
		log.Fatal(err)
	}

	runBrowser(selectedProfile)
}
