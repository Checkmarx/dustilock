package registry

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const httpRequestTimeout = time.Second * 10
const userAgent = "dependency locker"

const npmRegistryUrl = "https://registry.npmjs.org"
const pypiRegistryUrl = "https://pypi.python.org/simple"

func IsPypiPackageAvailableForRegistration(packageName string) (bool, error) {
	url := fmt.Sprintf("%v/%v", pypiRegistryUrl, packageName)
	return isPackageAvailableForRegistration(url)
}

func IsNpmPackageAvailableForRegistration(packageName string) (bool, error) {
	url := fmt.Sprintf("%v/%v", npmRegistryUrl, packageName)
	return isPackageAvailableForRegistration(url)
}

func isPackageAvailableForRegistration(url string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), httpRequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	isRegistered := resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest
	return !isRegistered, nil
}
