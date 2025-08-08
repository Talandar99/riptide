package main

import (
	"fmt"
	"github.com/Talandar99/riptide/internal"
	"os"
	"path/filepath"
	"strings"
)

func riptide(scriptsList []internal.Script) {
	var programArgsWithoutFlags, flagWithArgs []string
	programArgsWithoutFlags, flagWithArgs = internal.SeparateArgumentsAndFlags(os.Args[1:])

	if len(flagWithArgs) > 0 {
		switch flagWithArgs[0] {
		case "-r":
			fmt.Println(flagWithArgs)
			if len(flagWithArgs) == 2 {
				fmt.Println("Running remotely at " + flagWithArgs[1])
				for _, programArg := range programArgsWithoutFlags {
					var scriptsThatExists = internal.GetScriptsThatExists(programArg, scriptsList)
					for _, scriptThatExist := range scriptsThatExists {
						internal.RunRemoteCommand(scriptThatExist, flagWithArgs[1])
					}
				}
				os.Exit(0)
			} else {
				fmt.Println("Please specify address after -r")
				os.Exit(0)
			}
		default:
			fmt.Println("Unknown Flag")
			os.Exit(0)
		}
	}

	for _, programArg := range programArgsWithoutFlags {
		var scriptsThatExists = internal.GetScriptsThatExists(programArg, scriptsList)
		for _, scriptThatExist := range scriptsThatExists {
			internal.RunCommand(scriptThatExist)
		}
	}
	os.Exit(0)
}

func main() {
	progName := filepath.Base(os.Args[0])
	scriptsList := internal.GetScriptsList()
	//fmt.Println(scriptsList)

	if strings.Contains(progName, "completion") {
		//completion(scriptsList)
	} else {
		riptide(scriptsList)
	}
}
