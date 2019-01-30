package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRole strings.Replacer
var strconvRole strconv.NumError

var basicRoleTraitMapping = map[string]func(*schema.RoleTrait, []string){}
var customRoleTraitMapping = map[string]func(*schema.RoleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Role", func(ctx jsonld.Context) (interface{}) {
		return NewRoleFromContext(ctx)
	})






	customRoleTraitMapping["http://schema.org/endDate"] = func(x *schema.RoleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EndDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EndDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EndDate = s
		}
	}

	customRoleTraitMapping["http://schema.org/startDate"] = func(x *schema.RoleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.StartDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.StartDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.StartDate = s
		}
	}

	customRoleTraitMapping["http://schema.org/namedPosition"] = func(x *schema.RoleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NamedPosition, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NamedPosition = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NamedPosition = s
		}
	}

	customRoleTraitMapping["http://schema.org/roleName"] = func(x *schema.RoleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RoleName, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RoleName = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RoleName = s
		}
	}
}

func NewRoleFromContext(ctx jsonld.Context) (x schema.Role) {
	x.Type = "http://schema.org/Role"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRoleTrait(ctx, &x.RoleTrait)
	MapCustomToRoleTrait(ctx, &x.RoleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRoleTrait(ctx jsonld.Context, x *schema.RoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRoleTrait(ctx jsonld.Context, x *schema.RoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}