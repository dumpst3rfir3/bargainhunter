package bargainhunter

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func Fetch(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
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

var priceRE = regexp.MustCompile(`"FinalPrice":([\d.]+),"Instock"`)

func ExtractPrice(page string) (float64, error) {
	matches := priceRE.FindStringSubmatch(page)
	if len(matches) != 2 {
		return 0, errors.New("can't find price")
	}
	priceStr := matches[1]
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse %q as number", priceStr)
	}
	return price, nil
}
