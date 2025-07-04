package main

import (
	"fmt"
	"github.com/Talandar99/riptide/internal"
	"os"
	"path/filepath"
	"strings"
)

func riptide(scriptsList [][]string) {
	fmt.Println(scriptsList)

	var programArgsWithoutFlags, flagWithArgs []string
	programArgsWithoutFlags, flagWithArgs = internal.SeparateArgumentsAndFlags(os.Args[1:])

	if len(flagWithArgs) > 0 {
		switch flagWithArgs[0] {
		case "-r":
			os.Exit(0)
		default:
			fmt.Println("Unknown Flag")
			os.Exit(0)
		}
	}
	for _, arg := range programArgsWithoutFlags {
		internal.RunCommand(arg)
	}
	os.Exit(0)
}

func main() {
	progName := filepath.Base(os.Args[0])
	scriptsList := internal.GetScriptsList()
	fmt.Println(scriptsList)

	if strings.Contains(progName, "completion") {
		//completion(scriptsList)
	} else {
		riptide(scriptsList)
	}
}
