package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func getDefinitions(url string) (urls []string, err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	matches := prefixRegex.FindAllStringSubmatch(string(b), -1)

	for _, m := range matches {
		urls = append(urls, m[1])
	}

	return
}

func keepContentTypeContains(_urls []string, contains string) (urls []string, err error) {
	for _, u := range _urls {
		var r *http.Response
		r, err = http.Get(u)
		if err != nil {
			return
		}

		if !strings.Contains(r.Header.Get("Content-Type"), contains) {
			continue
		}

		urls = append(urls, u)
	}

	return
}

func resolveDefinitions(url string) (urls []string, err error) {
	known := map[string]bool{}

	ch := make(chan string, 100)

	ch <- url

FOR:
	for {
		select {
		case u := <-ch:
			_, ok := known[u]
			if ok {
				continue
			}
			known[u] = true

			var us []string
			us, err = getDefinitions(u)
			if err != nil {
				return
			}

			for _, u := range us {
				ch <- u
			}
		default:
			break FOR
		}
	}

	for u, _ := range known {
		urls = append(urls, u)
	}

	return
}