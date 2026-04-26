package util

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func ValidateDomain(path string, httpMode bool) (string, error) {
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		if httpMode {
			path = "http://" + path
		} else {
			path = "https://" + path
		}
	}

	u, err := url.ParseRequestURI(path)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return path, nil
	}
	if err == nil {
		err = fmt.Errorf("\"%s\" is invalid url/ip", path)
	}
	return path, err
}

func ValidateConnection(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not reach %s: %w", url, err)
	}

	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected return code: %d, when testing connection to %s", res.StatusCode, url)
	}

	return nil
}
