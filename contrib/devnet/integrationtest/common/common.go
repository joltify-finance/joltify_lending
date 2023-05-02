package common

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(cmdStr string, parameters ...string) {

	cmd := exec.Command(cmdStr, parameters...)
	cmd.Env = os.Environ()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Printf("%s\n", m)
	}
	cmd.Wait()

}

func RunCommandWithOutput(cmdStr string, parameters ...string) (string, error) {
	var outb, errb bytes.Buffer
	cmd := exec.Command(cmdStr, parameters...)
	cmd.Env = os.Environ()
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	if err := cmd.Run(); err != nil {
		return errb.String(), err
	}
	return outb.String(), nil
}
