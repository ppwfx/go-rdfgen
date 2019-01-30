package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFollowAction strings.Replacer
var strconvFollowAction strconv.NumError

var basicFollowActionTraitMapping = map[string]func(*schema.FollowActionTrait, []string){}
var customFollowActionTraitMapping = map[string]func(*schema.FollowActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FollowAction", func(ctx jsonld.Context) (interface{}) {
		return NewFollowActionFromContext(ctx)
	})



	customFollowActionTraitMapping["http://schema.org/followee"] = func(x *schema.FollowActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Followee, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Followee = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Followee = s
		}
	}
}

func NewFollowActionFromContext(ctx jsonld.Context) (x schema.FollowAction) {
	x.Type = "http://schema.org/FollowAction"
	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFollowActionTrait(ctx, &x.FollowActionTrait)
	MapCustomToFollowActionTrait(ctx, &x.FollowActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFollowActionTrait(ctx jsonld.Context, x *schema.FollowActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFollowActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFollowActionTrait(ctx jsonld.Context, x *schema.FollowActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFollowActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}