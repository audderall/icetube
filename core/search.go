package core

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Search(term string) (string, error) {
	encoded_term := url.QueryEscape(term)
	url := fmt.Sprintf("https://youtube.com/results?search_query=%s", encoded_term)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-arch", "x86")
	req.Header.Set("sec-ch-ua-platform", "Linux")
	req.Header.Set("Referer", "https://www.youtube.com/")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
