package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMenuSection strings.Replacer
var strconvMenuSection strconv.NumError

var basicMenuSectionTraitMapping = map[string]func(*schema.MenuSectionTrait, []string){}
var customMenuSectionTraitMapping = map[string]func(*schema.MenuSectionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MenuSection", func(ctx jsonld.Context) (interface{}) {
		return NewMenuSectionFromContext(ctx)
	})




	customMenuSectionTraitMapping["http://schema.org/hasMenuItem"] = func(x *schema.MenuSectionTrait, ctx jsonld.Context, s string){
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

	customMenuSectionTraitMapping["http://schema.org/hasMenuSection"] = func(x *schema.MenuSectionTrait, ctx jsonld.Context, s string){
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

func NewMenuSectionFromContext(ctx jsonld.Context) (x schema.MenuSection) {
	x.Type = "http://schema.org/MenuSection"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMenuSectionTrait(ctx, &x.MenuSectionTrait)
	MapCustomToMenuSectionTrait(ctx, &x.MenuSectionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMenuSectionTrait(ctx jsonld.Context, x *schema.MenuSectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMenuSectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMenuSectionTrait(ctx jsonld.Context, x *schema.MenuSectionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMenuSectionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}