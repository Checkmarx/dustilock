package registry

import (
	"testing"
)

func TestIsPackageAvailableForRegistration(t *testing.T) {

	existingPythonPackages := []string{"requests", "vmwc", "NumPy", "Pandas"}
	for _, packageName := range existingPythonPackages {
		isAvailableForRegistration, _ := IsPypiPackageAvailableForRegistration(packageName)
		if isAvailableForRegistration {
			t.Errorf("expected \"%v\" to be registered", packageName)
		}
	}

	existingNpmPackages := []string{"angular", "react"}
	for _, packageName := range existingNpmPackages {
		isAvailableForRegistration, _ := IsPypiPackageAvailableForRegistration(packageName)
		if isAvailableForRegistration {
			t.Errorf("expected \"%v\" to be registered", packageName)
		}
	}

	missingPackages := []string{"qweqeqezxczaw", "qweqeqe-zxczaw"}
	for _, packageName := range missingPackages {
		isAvailableForRegistration, _ := IsPypiPackageAvailableForRegistration(packageName)
		if !isAvailableForRegistration {
			t.Errorf("expected \"%v\" to be missing", packageName)
		}

		isAvailableForRegistration, _ = IsNpmPackageAvailableForRegistration(packageName)
		if !isAvailableForRegistration {
			t.Errorf("expected \"%v\" to be missing", packageName)
		}
	}

}
