package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ScriptsPath []string `toml:"scripts_path"`
}

func read_config() Config {
	var config Config
	toml.DecodeFile("config.toml", &config)
	return config
}

func getScriptsList() [][]string {
	config := read_config()
	var scriptsWithPath [][]string

	for _, scriptDirectory := range config.ScriptsPath {
		lsOutput, _ := exec.Command("ls", scriptDirectory).CombinedOutput()

		for lsOutputElement := range strings.SplitSeq(string(lsOutput), "\n") {
			lsOutputElement = strings.TrimSpace(lsOutputElement)

			if lsOutputElement != "" {
				scriptsWithPath = append(scriptsWithPath, []string{scriptDirectory, lsOutputElement})
			}
		}

	}
	return scriptsWithPath
}

func alreadyCompleted(completion_options []string) bool {
	program_args := os.Args
	program_args[0] = filepath.Base(program_args[0])
	for _, program_arg := range program_args {
		if slices.Contains(completion_options, program_arg) {
			return true
		}
	}
	return false
}

func completion(scriptsList []string) {
	compLine := os.Getenv("COMP_LINE")

	if alreadyCompleted(scriptsList) {
		return
	}

	prefix := ""
	words := strings.Fields(compLine)
	if len(words) > 1 {
		prefix = words[len(words)-1]
	}

	foundAnyCompletions := false
	for _, completionOption := range scriptsList {
		if strings.HasPrefix(completionOption, prefix) {
			fmt.Println(completionOption)
			foundAnyCompletions = true
		}
	}

	if !foundAnyCompletions && prefix == "" {
		for _, opt := range scriptsList {
			fmt.Println(opt)
		}
	}
}

func removeFlagsFromList(list []string, flags []string) []string {
	var listWithoutFlags []string
	for _, listArg := range list {
		var containFlag = false
		for _, flagArg := range flags {
			if strings.Contains(listArg, flagArg) {
				containFlag = true
				break
			}
		}
		if !containFlag {
			listWithoutFlags = append(listWithoutFlags, listArg)
		}
	}
	return listWithoutFlags
}

func getFlagsFromlist(list []string) []string {
	var flagWithArgs []string

	for i, programArg := range list {
		fmt.Println(programArg)
		if strings.Contains(programArg, "-r") {
			flagWithArgs = append(flagWithArgs, programArg)
			if i+1 < len(list) {
				flagWithArgs = append(flagWithArgs, list[i+1])
			}
		}
	}
	return flagWithArgs
}

func runCommand(command string) {
	fmt.Println("Running: " + command)
	cmd := exec.Command(command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

//func runRemoteCommand(command string, address string) {
//	fmt.Println("Running: " + command)
//	config := read_config()
//
//	fullCommand := config.ScriptsPath + command
//
//	cmd := exec.Command("scp " + fullCommand + " " + address + "~/" + command)
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	cmd.Stdin = os.Stdin
//
//	err := cmd.Run()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	cmd = exec.Command(" " + fullCommand + " " + address + "~/" + command)
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	cmd.Stdin = os.Stdin
//
//	err = cmd.Run()
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func runRemotely(programArgs []string, flagWithArgs []string) {
//	if len(flagWithArgs) > 2 {
//		//for _, program := range programArgs {
//		//cmd := exec.Command("ls", path)
//		//}
//	} else {
//		fmt.Println("missing address ")
//		fmt.Println("try:")
//		fmt.Println("riptide some_script.sh -r xyz@xyz.xyz")
//		os.Exit(1)
//	}
//	os.Exit(0)
//}

//func riptide(scriptsList []string) {
//	fmt.Println(scriptsList)
//	programArgs := os.Args[1:]
//
//	var flagWithArgs = getFlagsFromlist(programArgs)
//	fmt.Println(flagWithArgs)
//
//	programArgs = removeFlagsFromList(programArgs, flagWithArgs)
//	fmt.Println(programArgs)
//
//	if len(flagWithArgs) > 0 {
//		switch flagWithArgs[0] {
//		case "-r":
//			runRemotely(programArgs, flagWithArgs)
//			os.Exit(0)
//		default:
//			fmt.Println("Unknown Flag")
//			os.Exit(0)
//		}
//	}
//	for _, arg := range programArgs {
//		runCommand(arg)
//	}
//	os.Exit(0)
//}

func main() {
	//progName := filepath.Base(os.Args[0])
	scriptsList := getScriptsList()
	fmt.Println(scriptsList)

	if strings.Contains(progName, "completion") {
		completion(scriptsList)
	} else {
		riptide(scriptsList)
	}
}
