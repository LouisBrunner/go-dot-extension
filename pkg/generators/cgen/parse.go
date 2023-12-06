package main

import (
	"fmt"
	"os"
	"regexp"
)

var nameFinderRegexp = regexp.MustCompile(`(?s)\* @name ([^\n]+).*?\ntypedef [^\n]+?\(\*([^)]+)`)

func findInterfaceFuncs(inputPath string) (map[string]string, error) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	matches := nameFinderRegexp.FindAllStringSubmatch(string(content), -1)
	if matches == nil {
		return nil, fmt.Errorf("no interface function found")
	}
	for _, match := range matches {
		if len(match) != 3 {
			return nil, fmt.Errorf("invalid match: %+v", match)
		}
		fn := match[2]
		name := match[1]
		result[fn] = name
	}

	return result, nil
}
