package analysis

import (
	"bufio"
	"fmt"
	"github.com/dustico/dusti-lock/dependencies"
	"github.com/dustico/dusti-lock/registry"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func AnalyzePythonRequirementsFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	packageNames := dependencies.ParsePythonRequirements(reader)
	result := false

	for _, pythonPackageName := range packageNames {
		availableForRegistration, err := registry.IsPypiPackageAvailableForRegistration(pythonPackageName)

		if err != nil {
			fmt.Println(err)
			return false, err
		}

		if availableForRegistration {
			_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("error - python package \"%s\" is available for public registration. %s", pythonPackageName, filePath))
			result = true
		}
	}

	return result, nil
}

func AnalyzePackagesJsonFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	packageNames, err := dependencies.ParsePackagesJsonFile(reader)
	if err != nil {
		return false, err
	}

	result := false

	for _, packageName := range packageNames {
		availableForRegistration, err := registry.IsNpmPackageAvailableForRegistration(packageName)

		if err != nil {
			fmt.Println(err)
			return false, err
		}

		if availableForRegistration {
			_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf("error - npm package \"%s\" is available for public registration. %s", packageName, filePath))
			result = true
		}
	}

	return result, nil
}

func AnalyzeDirectoryRecursive(workingDir string, excludedDirectories map[string]bool) (bool, error) {
	hasAnyPackageAvailableForRegistration := false

	err := filepath.Walk(workingDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if excludedDirectories != nil {
			_, isExcluded := excludedDirectories[fileInfo.Name()]
			if fileInfo.IsDir() && isExcluded {
				return filepath.SkipDir
			}
		}

		fileName := fileInfo.Name()
		if fileName == "package.json" {
			result, err := AnalyzePackagesJsonFile(path)
			if result {
				hasAnyPackageAvailableForRegistration = true
			}

			if err != nil {
				fmt.Println(err)
			}
			return err
		}

		if fileName == "requirements.txt" {
			result, err := AnalyzePythonRequirementsFile(path)
			if result {
				hasAnyPackageAvailableForRegistration = true
			}

			if err != nil {
				fmt.Println(err)
			}
			return err
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return hasAnyPackageAvailableForRegistration, nil
}

func AnalyzeDirectory(workingDir string) (bool, error) {
	hasAnyPackageAvailableForRegistration := false

	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return false, err
	}

	for _, fileInfo := range files {
		fileName := fileInfo.Name()
		filePath := path.Join(workingDir, fileName)

		if fileName == "package.json" {
			result, err := AnalyzePackagesJsonFile(filePath)
			if result {
				hasAnyPackageAvailableForRegistration = true
			}

			if err != nil {
				return false, err
			}
		}

		if fileName == "requirements.txt" {
			result, err := AnalyzePythonRequirementsFile(filePath)
			if result {
				hasAnyPackageAvailableForRegistration = true
			}

			if err != nil {
				return false, err
			}
		}
	}

	return hasAnyPackageAvailableForRegistration, nil
}
