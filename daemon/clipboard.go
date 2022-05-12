package daemon

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetValue() string {
	cmd := exec.Command("pbpaste")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
    if err != nil {
		log.Fatal(err)
        os.Exit(1)
	}

	block := ""
	scanner := bufio.NewScanner(bufio.NewReader(stdout))
	for scanner.Scan() {
		if block != "" {
			block += "\n"
		}

		block += scanner.Text()
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error Waiting process: " + err.Error())
	}

	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error killing process: " + err.Error())
	}

	return block
}