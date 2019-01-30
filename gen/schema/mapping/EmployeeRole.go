package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmployeeRole strings.Replacer
var strconvEmployeeRole strconv.NumError

var basicEmployeeRoleTraitMapping = map[string]func(*schema.EmployeeRoleTrait, []string){}
var customEmployeeRoleTraitMapping = map[string]func(*schema.EmployeeRoleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EmployeeRole", func(ctx jsonld.Context) (interface{}) {
		return NewEmployeeRoleFromContext(ctx)
	})



	basicEmployeeRoleTraitMapping["http://schema.org/salaryCurrency"] = func(x *schema.EmployeeRoleTrait, s []string) {
		x.SalaryCurrency = s[0]
	}


	customEmployeeRoleTraitMapping["http://schema.org/baseSalary"] = func(x *schema.EmployeeRoleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BaseSalary, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BaseSalary = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BaseSalary = s
		}
	}
}

func NewEmployeeRoleFromContext(ctx jsonld.Context) (x schema.EmployeeRole) {
	x.Type = "http://schema.org/EmployeeRole"
	MapBasicToOrganizationRoleTrait(ctx, &x.OrganizationRoleTrait)
	MapCustomToOrganizationRoleTrait(ctx, &x.OrganizationRoleTrait)

	MapBasicToRoleTrait(ctx, &x.RoleTrait)
	MapCustomToRoleTrait(ctx, &x.RoleTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEmployeeRoleTrait(ctx, &x.EmployeeRoleTrait)
	MapCustomToEmployeeRoleTrait(ctx, &x.EmployeeRoleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmployeeRoleTrait(ctx jsonld.Context, x *schema.EmployeeRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmployeeRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmployeeRoleTrait(ctx jsonld.Context, x *schema.EmployeeRoleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmployeeRoleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}