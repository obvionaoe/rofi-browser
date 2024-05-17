package cmd

import (
	"github.com/Wing924/shellwords"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const (
	commonBrowserArgs    = "--new-instance"
	profileManagerOption = "Launch Profile Manager"
	profileManagerArg    = "--ProfileManager"
	profilePickerArg     = "-P"
)

func getProfiles(path string) ([]string, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	var profiles []string

	for _, sec := range cfg.Sections() {
		if strings.Contains(sec.Name(), "Profile") {
			profiles = append(profiles, sec.Key("Name").String())
		}
	}

	return profiles, nil
}

func runBrowser(profile string) {
	var args string

	if profile == profileManagerOption {
		args = strings.Join([]string{commonBrowserArgs, profileManagerArg}, " ")
	} else {
		args = strings.Join([]string{commonBrowserArgs, profilePickerArg, profile}, " ")
	}

	escapedArgs, err := shellwords.Split(args)
	if err != nil {
		log.Fatal(err)
	}

	executable, err := exec.LookPath(browser)
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Exec(executable, escapedArgs, os.Environ())
	if err != nil {
		log.Fatal(err)
	}
}
