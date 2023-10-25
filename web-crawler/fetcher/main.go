package fetcher

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func GetBody(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error while making http request: " + err.Error())
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error while making http request: " + err.Error())
	}

	return string(body), nil
}

func GetLinksFromBody(body string) ([]string, error) {
	// extract urls from the body
	// href="<url>"
	urlExpression := `(href|src)="(https?://[A-Za-z0-9_.\-~/?=#&]*)"`
	urlRegex, err := regexp.Compile(urlExpression)
	if err != nil {
		return nil, err
	}

	ans := make([]string, 0, 10)
	relatedUrls := urlRegex.FindAllStringSubmatch(body, -1)
	for _, relatedurl := range relatedUrls {
		ans = append(ans, relatedurl[2])
	}

	return ans, nil
}

func GetRelatedLinks(url string) []string {
	body, err := GetBody(url)
	if err != nil {
		log.Default().Fatal(err.Error())
	}

	ans, err := GetLinksFromBody(body)
	if err != nil {
		return nil
	}

	return ans
}
