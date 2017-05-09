package command

import (
	"log"
	"os"
	"os/exec"
)

func Run(cmd string) error {
	log.Println("exec: ", cmd)

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
