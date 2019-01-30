package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMenu strings.Replacer
var strconvMenu strconv.NumError

var basicMenuTraitMapping = map[string]func(*schema.MenuTrait, []string){}
var customMenuTraitMapping = map[string]func(*schema.MenuTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Menu", func(ctx jsonld.Context) (interface{}) {
		return NewMenuFromContext(ctx)
	})




	customMenuTraitMapping["http://schema.org/hasMenuItem"] = func(x *schema.MenuTrait, ctx jsonld.Context, s string){
		var y schema.MenuItem
		if strings.HasPrefix(s, "_:") {
			y = NewMenuItemFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMenuItem()
			y.Id = s
		}

		x.HasMenuItem = &y

		return
	}

	customMenuTraitMapping["http://schema.org/hasMenuSection"] = func(x *schema.MenuTrait, ctx jsonld.Context, s string){
		var y schema.MenuSection
		if strings.HasPrefix(s, "_:") {
			y = NewMenuSectionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMenuSection()
			y.Id = s
		}

		x.HasMenuSection = &y

		return
	}
}

func NewMenuFromContext(ctx jsonld.Context) (x schema.Menu) {
	x.Type = "http://schema.org/Menu"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMenuTrait(ctx, &x.MenuTrait)
	MapCustomToMenuTrait(ctx, &x.MenuTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMenuTrait(ctx jsonld.Context, x *schema.MenuTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMenuTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMenuTrait(ctx jsonld.Context, x *schema.MenuTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMenuTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}