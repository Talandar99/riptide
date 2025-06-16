package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type Config struct {
	ScriptsPath []string `toml:"scripts_path"`
}

func read_config() Config {
	var config Config
	toml.DecodeFile("config.toml", &config)
	return config
}

func getScriptsList() []string {
	config := read_config()
	var scripts []string
	for _, path := range config.ScriptsPath {
		cmd := exec.Command("ls", path)
		output, _ := cmd.CombinedOutput()
		for line := range strings.SplitSeq(string(output), "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				scripts = append(scripts, line)
			}
		}
	}
	return scripts
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

func riptide(scriptsList []string) {
	fmt.Println(scriptsList)
	programArgs := os.Args[1:]

	var flagWithArgs = getFlagsFromlist(programArgs)
	fmt.Println(flagWithArgs)

	programArgs = removeFlagsFromList(programArgs, flagWithArgs)
	fmt.Println(programArgs)
}

func main() {
	progName := filepath.Base(os.Args[0])
	scriptsList := getScriptsList()
	if strings.Contains(progName, "completion") {
		completion(scriptsList)
	} else {
		riptide(scriptsList)
	}
}
