package cmd

import (
	"github.com/Wing924/shellwords"
	"io"
	"log"
	"os/exec"
	"strings"
)

func runRofi(profiles []string) (string, error) {
	rofiArgs, err := shellwords.Split(rofiCmd)
	if err != nil {
		log.Fatal(err)
	}

	rofi := exec.Command(rofiArgs[0], rofiArgs[1:]...)
	stdin, err := rofi.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer func(stdin io.WriteCloser) {
			err := stdin.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(stdin)
		for _, s := range profiles {
			_, err := io.WriteString(stdin, s+"\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	out, err := rofi.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out)), nil
}
