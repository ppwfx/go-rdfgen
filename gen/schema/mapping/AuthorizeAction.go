package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAuthorizeAction strings.Replacer
var strconvAuthorizeAction strconv.NumError

var basicAuthorizeActionTraitMapping = map[string]func(*schema.AuthorizeActionTrait, []string){}
var customAuthorizeActionTraitMapping = map[string]func(*schema.AuthorizeActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AuthorizeAction", func(ctx jsonld.Context) (interface{}) {
		return NewAuthorizeActionFromContext(ctx)
	})



	customAuthorizeActionTraitMapping["http://schema.org/recipient"] = func(x *schema.AuthorizeActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Recipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Recipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Recipient = s
		}
	}
}

func NewAuthorizeActionFromContext(ctx jsonld.Context) (x schema.AuthorizeAction) {
	x.Type = "http://schema.org/AuthorizeAction"
	MapBasicToAllocateActionTrait(ctx, &x.AllocateActionTrait)
	MapCustomToAllocateActionTrait(ctx, &x.AllocateActionTrait)

	MapBasicToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)
	MapCustomToOrganizeActionTrait(ctx, &x.OrganizeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAuthorizeActionTrait(ctx, &x.AuthorizeActionTrait)
	MapCustomToAuthorizeActionTrait(ctx, &x.AuthorizeActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAuthorizeActionTrait(ctx jsonld.Context, x *schema.AuthorizeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAuthorizeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAuthorizeActionTrait(ctx jsonld.Context, x *schema.AuthorizeActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAuthorizeActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}