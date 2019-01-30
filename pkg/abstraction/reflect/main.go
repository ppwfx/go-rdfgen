package main

import (
	"log"
	_ "github.com/21stio/go-rdfgen/gen/schema/mapping"
	"github.com/21stio/go-rdfgen/gen/schema"
	"encoding/json"
)

var j1 = []byte(`
{
	"@context": [
		"http://schema.org",
		{
			"mkay": "@type",
			"yoyo": "additionalName"
		}
	],
	"@id": "http://yoyo.com/Peter",
	"mkay": "http://schema.org/Person",
	"yoyo": "Peter"
}`)

var j2 = []byte(`
{
	"@context":"http://schema.org",
	"@id": "http://yoyo.com/Peter",
	"@type": "http://schema.org/Event",
	"inLanguage": "EN"
}`)

var j3 = []byte(`
{
  "@context": "http://schema.org",
  "@id": "http://schema.org/Peter",
  "@type": "Person",
  "address": {
    "@type": "PostalAddress",
    "addressLocality": "Seattle",
    "addressRegion": "WA",
    "postalCode": "98052",
    "streetAddress": "20341 Whitworth Institute 405 N. Whitworth"
  },
  "colleague": [
    "http://www.xyz.edu/students/alicejones.html",
    "http://www.xyz.edu/students/bobsmith.html"
  ],
  "email": "mailto:jane-doe@xyz.edu",
  "image": "/janedoe.jpg",
  "jobTitle": "Professor",
  "name": "Jane Doe",
  "telephone": "(425) 123-4567",
  "url": "http://www.janedoe.com"
}`)

var j4 = []byte(`{
  "@context": "http://schema.org/",
  "@type": "Person",
  "@id": "http://schema.org/Peter",
  "name": "Christopher Froome",
  "sponsor":
  {
    "@type": "Organization",
    "name": "Sky",
    "url": "http://www.skysports.com/"
  },
  "address": {
	"@type": "PostalAddress",
	"addressLocality": "Seattle",
	"addressRegion": "WA",
	"postalCode": "98052",
	"streetAddress": "20341 Whitworth Institute 405 N. Whitworth"
  },
  "colleague": [
	"http://www.xyz.edu/students/alicejones.html",
	"http://www.xyz.edu/students/bobsmith.html"
  ]
}`)

var j5 = []byte(`{
  "@context": [
    "http://schema.org/"
  ],
  "@type": "Person",
  "@id": "http://PersonAbc",
  "name": "Christopher Froome",
  "sponsor":
  {
    "@type": "Organization",
    "name": "Sky",
    "url": "http://www.skysports.com/"
  },
  "yo": {
    "@type": "Game",
    "offers":{
      "@type":"Offer",
      "priceCurrency":"$",
      "price":"17.99",
      "availableAtOrFrom":"http://monopoly-2/en_US/shop/where-to-buy.cfm?brand_guid=DAD28866-1C43-11DD-BD0B-0800200C9A66&prodName=Monopoly%20Game"
    }
  }
}`)

type PersonMarsh struct {
	Context interface{} `json:"@context"`
	schema.Person
	Jo      interface{}
}

func MarshalPerson(p schema.Person) (b []byte, err error) {
	//d := map[string]interface{}{}
	//
	//m := structs.Map(v)

	//for _, v0 := range m {
	//	v0, ok := v0.(map[string]interface{})
	//	if !ok {
	//
	//	}
	//
	//	//if k0 == "AdditionalTrait" {
	//	//	for k1, v1 := range v0["AdditionalTrait"]["Additional"] {
	//	//		a[k] = v
	//	//	}
	//	//}
	//
	//	for k1, v1 := range v0 {
	//		d[k1] = v1
	//	}
	//}

	pm := PersonMarsh{}
	pm.Person = p
	pm.Context = "yo"
	//pm.Jo = map[string]interface{}{}
	//pm.Jo["asds"] = "asd"

	b, err = json.MarshalIndent(pm, "", "   ")

	println(string(b))

	return
}

func main() {
	err := func() (err error) {
		//v, err := jsonld.Unmarshal(j5)
		//if err != nil {
		//	return
		//}

		p := schema.NewPerson()
		p.Id = "http://PersonAbc"
		p.Name = "Christopher Froome"

		sponsor := schema.NewOrganization()
		sponsor.Name = "Sky"
		sponsor.Url = "http://www.skysports.com/"

		g := schema.NewGame()

		offer := schema.NewOffer()
		offer.PriceCurrency = "$"
		offer.Price = 17.99

		place := schema.NewPlace()
		place.Id = "http://monopoly-2/en_US/shop/where-to-buy.cfm?brand_guid=DAD28866-1C43-11DD-BD0B-0800200C9A66&prodName=Monopoly%20Game"
		offer.AvailableAtOrFrom = &place

		g.Offers = &offer

		p.Sponsor = sponsor

		//b, err := json.MarshalIndent(p, "", "   ")
		//println(string(b))

		MarshalPerson(p)

		return
	}()
	if err != nil {
		log.Fatal(err)
	}

}
