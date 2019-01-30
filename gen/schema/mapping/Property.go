package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsProperty strings.Replacer
var strconvProperty strconv.NumError

var basicPropertyTraitMapping = map[string]func(*schema.PropertyTrait, []string){}
var customPropertyTraitMapping = map[string]func(*schema.PropertyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Property", func(ctx jsonld.Context) (interface{}) {
		return NewPropertyFromContext(ctx)
	})






	customPropertyTraitMapping["http://schema.org/domainIncludes"] = func(x *schema.PropertyTrait, ctx jsonld.Context, s string){
		var y schema.Class
		if strings.HasPrefix(s, "_:") {
			y = NewClassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewClass()
			y.Id = s
		}

		x.DomainIncludes = &y

		return
	}

	customPropertyTraitMapping["http://schema.org/supersededBy"] = func(x *schema.PropertyTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SupersededBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SupersededBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SupersededBy = s
		}
	}

	customPropertyTraitMapping["http://schema.org/inverseOf"] = func(x *schema.PropertyTrait, ctx jsonld.Context, s string){
		var y schema.Property
		if strings.HasPrefix(s, "_:") {
			y = NewPropertyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewProperty()
			y.Id = s
		}

		x.InverseOf = &y

		return
	}

	customPropertyTraitMapping["http://schema.org/rangeIncludes"] = func(x *schema.PropertyTrait, ctx jsonld.Context, s string){
		var y schema.Class
		if strings.HasPrefix(s, "_:") {
			y = NewClassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewClass()
			y.Id = s
		}

		x.RangeIncludes = &y

		return
	}
}

func NewPropertyFromContext(ctx jsonld.Context) (x schema.Property) {
	x.Type = "http://schema.org/Property"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPropertyTrait(ctx, &x.PropertyTrait)
	MapCustomToPropertyTrait(ctx, &x.PropertyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPropertyTrait(ctx jsonld.Context, x *schema.PropertyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPropertyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPropertyTrait(ctx jsonld.Context, x *schema.PropertyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPropertyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}