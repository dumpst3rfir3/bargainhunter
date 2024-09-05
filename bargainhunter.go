package bargainhunter

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP response status: %s", resp.Status)
	}

	bodytext, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodytext), nil
}
