package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2020/08/gameconsole"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}

	gc := gameconsole.NewGameConsole()
	if err := gc.LoadProgram(string(input)); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed loading program:", err)
	}

	gc.RunAndStopOnReexecution()
	fmt.Println("acc at loop:", gc.Acc())
}
