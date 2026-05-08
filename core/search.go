package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Search(term string) {
	encoded_term := url.QueryEscape(term)
	url := fmt.Sprintf("https://youtube.com/?search_query=%s", encoded_term)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
