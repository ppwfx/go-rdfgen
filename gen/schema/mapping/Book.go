package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBook strings.Replacer
var strconvBook strconv.NumError

var basicBookTraitMapping = map[string]func(*schema.BookTrait, []string){}
var customBookTraitMapping = map[string]func(*schema.BookTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Book", func(ctx jsonld.Context) (interface{}) {
		return NewBookFromContext(ctx)
	})


	basicBookTraitMapping["http://schema.org/abridged"] = func(x *schema.BookTrait, s []string) {
		var err error
		x.Abridged, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicBookTraitMapping["http://schema.org/bookEdition"] = func(x *schema.BookTrait, s []string) {
		x.BookEdition = s[0]
	}




	basicBookTraitMapping["http://schema.org/isbn"] = func(x *schema.BookTrait, s []string) {
		x.Isbn = s[0]
	}



	customBookTraitMapping["http://schema.org/bookFormat"] = func(x *schema.BookTrait, ctx jsonld.Context, s string){
		var y schema.BookFormatType
		if strings.HasPrefix(s, "_:") {
			y = NewBookFormatTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBookFormatType()
			y.Id = s
		}

		x.BookFormat = &y

		return
	}

	customBookTraitMapping["http://schema.org/illustrator"] = func(x *schema.BookTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Illustrator = &y

		return
	}

	customBookTraitMapping["http://schema.org/numberOfPages"] = func(x *schema.BookTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.NumberOfPages = &y

		return
	}
}

func NewBookFromContext(ctx jsonld.Context) (x schema.Book) {
	x.Type = "http://schema.org/Book"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBookTrait(ctx, &x.BookTrait)
	MapCustomToBookTrait(ctx, &x.BookTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBookTrait(ctx jsonld.Context, x *schema.BookTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBookTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBookTrait(ctx jsonld.Context, x *schema.BookTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBookTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}