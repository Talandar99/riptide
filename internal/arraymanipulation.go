package internal

import (
	"fmt"
	"strings"
)

func SeparateArgumentsAndFlags(list []string) ([]string, []string) {
	var flagWithParameters []string
	var ProgArgWithoutFlags []string

	for i, programArg := range list {
		fmt.Println(programArg)
		if strings.Contains(programArg, "-r") {
			flagWithParameters = append(flagWithParameters, programArg)
			if i+1 < len(list) {
				flagWithParameters = append(flagWithParameters, list[i+1])
			}
		}
	}
	for _, listMember := range list {
		var containFlag = false
		for _, flagMember := range flagWithParameters {
			if strings.Contains(listMember, flagMember) {
				containFlag = true
				break
			}
		}
		if !containFlag {
			ProgArgWithoutFlags = append(ProgArgWithoutFlags, listMember)
		}
	}
	return ProgArgWithoutFlags, flagWithParameters

}
