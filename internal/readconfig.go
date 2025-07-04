package internal

import (
	"os/exec"
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

func GetScriptsList() [][]string {
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
