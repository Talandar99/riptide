package internal

import (
	"github.com/BurntSushi/toml"
	"os/exec"
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

func GetScriptsList() []Script {

	config := read_config()
	var scripts []Script

	for _, scriptDirectory := range config.ScriptsPath {
		lsOutput, _ := exec.Command("ls", scriptDirectory).CombinedOutput()

		for lsOutputElement := range strings.SplitSeq(string(lsOutput), "\n") {
			lsOutputElement = strings.TrimSpace(lsOutputElement)

			if lsOutputElement != "" {
				//scriptsWithPath = append(scriptsWithPath, []string{scriptDirectory, lsOutputElement})
				scripts = append(scripts, Script{Name: lsOutputElement, Path: scriptDirectory})
			}
		}

	}
	return scripts
}
