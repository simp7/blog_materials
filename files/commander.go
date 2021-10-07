package files

import (
	"fmt"
	"os/exec"
)

type Commander struct {
	file string
}

func New(file string) Commander {
	return Commander{file}
}

func (c Commander) Do(arg string) {
	cmd := exec.Command("python3", c.file, arg)
	result, err := cmd.Output()
	fmt.Println(err)
	fmt.Printf("Output : %s\n", string(result))
}
