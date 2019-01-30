package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLinkRole strings.Replacer
var strconvLinkRole strconv.NumError

var basicLinkRoleTraitMapping = map[string]func(*schema.LinkRoleTrait, []string){}
var customLinkRoleTraitMapping = map[string]func(*schema.LinkRoleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LinkRole", func(ctx jsonld.Context) (interface{}) {
		return NewLinkRoleFromContext(ctx)
	})



	basicLinkRoleTraitMapping["http://schema.org/linkRelationship"] = func(x *schema.LinkRoleTrait, s []string) {
		x.LinkRelationship = s[0]
	}


	customLinkRoleTraitMapping["http://schema.org/inLanguage"] = func(x *schema.LinkRoleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InLanguage = s
		}
	}
}

func NewLinkRoleFromContext(ctx jsonld.Context) (x schema.LinkRole) {
	x.Type = "http://schema.org/LinkRole"
	MapBasicToRoleTrait(ctx, &x.RoleTrait)
	MapCustomToRoleTrait(ctx, &x.RoleTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLinkRoleTrait(ctx, &x.LinkRoleTrait)
	MapCustomToLinkRoleTrait(ctx, &x.LinkRoleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLinkRoleTrait(ctx jsonld.Context, x *schema.LinkRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLinkRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLinkRoleTrait(ctx jsonld.Context, x *schema.LinkRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLinkRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}