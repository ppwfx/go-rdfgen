package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
)

type metaTrait struct {
	Id string `json:"@id,omitempty"`
	Type string `json:"@type,omitempty"`
}

type A struct {
	metaTrait
	AString string `json:"aString,omitempty"`
}

func NewA() (b A) {
	b.Type = "http://schema.org/A"

	return
}

type B struct {
	metaTrait
	BInt int `json:"aInt,omitempty"`
}

func NewB() (b B) {
	b.Type = "http://schema.org/B"

	return
}

type Review struct {
	metaTrait
	ItemReviewed interface{} `json:"itemReviewed,omitempty"`
}

func NewReview() (r Review) {
	r.Type = "http://schema.org/Review"

	return
}

var types = map[string]interface{}{}

func init() {
	types["http://schema.org/Review"] = Review{}
	types["http://schema.org/A"] = A{}
	types["http://schema.org/B"] = B{}
}

func main()  {
	err := func () (err error) {
		a := NewA()
		a.AString = "foo"

		r := NewReview()
		r.ItemReviewed = a

		b, err := json.Marshal(r)
		if err != nil {
			return
		}

		mt := &metaTrait{}
		r2 := Review{}
		r2.ItemReviewed = &mt
		err = json.Unmarshal(b, &r2)
		if err != nil {
			return
		}

		sth := types[mt.Type]
		r2.ItemReviewed = &sth
		err = json.Unmarshal(b, &r2)
		if err != nil {
			return
		}

		spew.Dump(sth)

		return
	}()
	if err != nil {
		log.Fatal(err)
	}
}