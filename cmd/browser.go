package cmd

import (
	"fmt"
	"github.com/Wing924/shellwords"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const args = "--new-instance -P"

func getProfiles(path string) ([]string, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Println("test")
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
	args := strings.Join([]string{args, profile}, " ")
	escapedArgs, err := shellwords.Split(args)
	if err != nil {
		return
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
