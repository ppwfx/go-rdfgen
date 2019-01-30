package utils

import (
	"net/http"
	"github.com/21stio/go-rdfgen/pkg/types"
	"os"
	"net/url"
	"encoding/gob"
	"io/ioutil"
	"strings"
)

func Get(u string) (r types.Response, err error) {
	var file *os.File

	p := "./cache/" + escapePath(u)

	_, err = os.Stat(p)
	if err == nil {
		file, err = os.Open(p)
		if err != nil {
			return
		}
		defer file.Close()

		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&r)
		if err != nil {
			return
		}

		return
	}

	rsp, err := http.Get(u)
	if err != nil {
		return
	}

	r.Body, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	r.Header = rsp.Header
	file, err = os.Create(p)
	if err != nil {
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(r)
	if err != nil {
		return
	}

	return
}

func escapePath(s string) (string) {
	s = url.PathEscape(s)

	// some fs match case-insensitive
	s = strings.Replace(s, "A", "aaa", -1)
	s = strings.Replace(s, "B", "bbb", -1)
	s = strings.Replace(s, "C", "ccc", -1)
	s = strings.Replace(s, "D", "ddd", -1)
	s = strings.Replace(s, "E", "eee", -1)
	s = strings.Replace(s, "F", "fff", -1)
	s = strings.Replace(s, "G", "ggg", -1)
	s = strings.Replace(s, "H", "hhh", -1)
	s = strings.Replace(s, "I", "iii", -1)
	s = strings.Replace(s, "J", "jjj", -1)
	s = strings.Replace(s, "K", "kkk", -1)
	s = strings.Replace(s, "L", "lll", -1)
	s = strings.Replace(s, "M", "mmm", -1)
	s = strings.Replace(s, "N", "nnn", -1)
	s = strings.Replace(s, "O", "ooo", -1)
	s = strings.Replace(s, "P", "ppp", -1)
	s = strings.Replace(s, "Q", "qqq", -1)
	s = strings.Replace(s, "R", "rrr", -1)
	s = strings.Replace(s, "S", "sss", -1)
	s = strings.Replace(s, "T", "ttt", -1)
	s = strings.Replace(s, "U", "uuu", -1)
	s = strings.Replace(s, "V", "vvv", -1)
	s = strings.Replace(s, "W", "www", -1)
	s = strings.Replace(s, "X", "xxx", -1)
	s = strings.Replace(s, "Y", "yyy", -1)
	s = strings.Replace(s, "Z", "zzz", -1)

	return s
}