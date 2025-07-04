package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

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
