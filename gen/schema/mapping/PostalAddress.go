package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPostalAddress strings.Replacer
var strconvPostalAddress strconv.NumError

var basicPostalAddressTraitMapping = map[string]func(*schema.PostalAddressTrait, []string){}
var customPostalAddressTraitMapping = map[string]func(*schema.PostalAddressTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PostalAddress", func(ctx jsonld.Context) (interface{}) {
		return NewPostalAddressFromContext(ctx)
	})



	basicPostalAddressTraitMapping["http://schema.org/postalCode"] = func(x *schema.PostalAddressTrait, s []string) {
		x.PostalCode = s[0]
	}


	basicPostalAddressTraitMapping["http://schema.org/addressLocality"] = func(x *schema.PostalAddressTrait, s []string) {
		x.AddressLocality = s[0]
	}


	basicPostalAddressTraitMapping["http://schema.org/addressRegion"] = func(x *schema.PostalAddressTrait, s []string) {
		x.AddressRegion = s[0]
	}


	basicPostalAddressTraitMapping["http://schema.org/postOfficeBoxNumber"] = func(x *schema.PostalAddressTrait, s []string) {
		x.PostOfficeBoxNumber = s[0]
	}


	basicPostalAddressTraitMapping["http://schema.org/streetAddress"] = func(x *schema.PostalAddressTrait, s []string) {
		x.StreetAddress = s[0]
	}


	customPostalAddressTraitMapping["http://schema.org/addressCountry"] = func(x *schema.PostalAddressTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AddressCountry, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AddressCountry = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AddressCountry = s
		}
	}
}

func NewPostalAddressFromContext(ctx jsonld.Context) (x schema.PostalAddress) {
	x.Type = "http://schema.org/PostalAddress"
	MapBasicToContactPointTrait(ctx, &x.ContactPointTrait)
	MapCustomToContactPointTrait(ctx, &x.ContactPointTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPostalAddressTrait(ctx, &x.PostalAddressTrait)
	MapCustomToPostalAddressTrait(ctx, &x.PostalAddressTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPostalAddressTrait(ctx jsonld.Context, x *schema.PostalAddressTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPostalAddressTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPostalAddressTrait(ctx jsonld.Context, x *schema.PostalAddressTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPostalAddressTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}