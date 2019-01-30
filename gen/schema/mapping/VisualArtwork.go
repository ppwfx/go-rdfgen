package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVisualArtwork strings.Replacer
var strconvVisualArtwork strconv.NumError

var basicVisualArtworkTraitMapping = map[string]func(*schema.VisualArtworkTrait, []string){}
var customVisualArtworkTraitMapping = map[string]func(*schema.VisualArtworkTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VisualArtwork", func(ctx jsonld.Context) (interface{}) {
		return NewVisualArtworkFromContext(ctx)
	})















	customVisualArtworkTraitMapping["http://schema.org/height"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Height, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Height = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Height = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/width"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Width, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Width = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Width = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/depth"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Depth, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Depth = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Depth = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/artEdition"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ArtEdition, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ArtEdition = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ArtEdition = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/artMedium"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ArtMedium, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ArtMedium = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ArtMedium = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/artform"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Artform, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Artform = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Artform = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/artist"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Artist = &y

		return
	}

	customVisualArtworkTraitMapping["http://schema.org/artworkSurface"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ArtworkSurface, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ArtworkSurface = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ArtworkSurface = s
		}
	}

	customVisualArtworkTraitMapping["http://schema.org/colorist"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Colorist = &y

		return
	}

	customVisualArtworkTraitMapping["http://schema.org/inker"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Inker = &y

		return
	}

	customVisualArtworkTraitMapping["http://schema.org/letterer"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Letterer = &y

		return
	}

	customVisualArtworkTraitMapping["http://schema.org/penciler"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Penciler = &y

		return
	}

	customVisualArtworkTraitMapping["http://schema.org/surface"] = func(x *schema.VisualArtworkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Surface, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Surface = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Surface = s
		}
	}
}

func NewVisualArtworkFromContext(ctx jsonld.Context) (x schema.VisualArtwork) {
	x.Type = "http://schema.org/VisualArtwork"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVisualArtworkTrait(ctx, &x.VisualArtworkTrait)
	MapCustomToVisualArtworkTrait(ctx, &x.VisualArtworkTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVisualArtworkTrait(ctx jsonld.Context, x *schema.VisualArtworkTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVisualArtworkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVisualArtworkTrait(ctx jsonld.Context, x *schema.VisualArtworkTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVisualArtworkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}