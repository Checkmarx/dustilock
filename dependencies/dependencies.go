package dependencies

import (
	"bufio"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func ParsePythonRequirements(reader *bufio.Reader) []string {
	packageNamesSet := map[string]bool{}

	for {
		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		line := string(lineBytes)
		line = strings.TrimSpace(line)

		re := regexp.MustCompile(`[#&]+egg=([a-zA-Z0-9_\-.]+)`)
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			packageName := strings.ToLower(match[1])
			packageNamesSet[packageName] = true
			continue
		}

		line = strings.Split(line, "#")[0]

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "--requirement") {
			continue
		}

		if strings.HasPrefix(line, "-r") {
			continue
		}

		if strings.Contains(line, "://") {
			continue
		}

		re = regexp.MustCompile(`^([a-zA-Z0-9_\-.]+)`)
		match = re.FindStringSubmatch(line)
		if len(match) > 0 {
			packageName := strings.ToLower(match[1])
			packageNamesSet[packageName] = true
			continue
		}

	}

	packageNames := []string{}
	for k := range packageNamesSet {
		packageNames = append(packageNames, k)
	}
	return packageNames
}

func ParsePackagesJsonFile(reader *bufio.Reader) ([]string, error) {
	packageNamesSet := map[string]bool{}

	d := json.NewDecoder(reader)
	t := struct {
		Dependencies    *map[string]interface{} `json:"dependencies"`
		DevDependencies *map[string]interface{} `json:"devDependencies"`
	}{}

	err := d.Decode(&t)
	if err != nil {
		return nil, err
	}

	processPackageName := func(dict *map[string]interface{}, npmPackageName string) {
		if strings.HasPrefix(npmPackageName, "@") {
			return
		}

		value, _ := (*dict)[npmPackageName]
		version := fmt.Sprintf("%v", value)
		version = strings.ToLower(version)

		if strings.HasPrefix(version, "npm:") {
			return
		}

		if strings.Contains(version, "://") {
			return
		}

		packageNamesSet[npmPackageName] = true
	}

	if t.Dependencies != nil {
		for npmPackageName := range *t.Dependencies {
			processPackageName(t.Dependencies, npmPackageName)
		}
	}

	if t.DevDependencies != nil {
		for npmPackageName := range *t.DevDependencies {
			processPackageName(t.DevDependencies, npmPackageName)
		}
	}

	packageNames := []string{}
	for k := range packageNamesSet {
		packageNames = append(packageNames, k)
	}
	return packageNames, err
}
