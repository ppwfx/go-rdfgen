package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRsvpAction strings.Replacer
var strconvRsvpAction strconv.NumError

var basicRsvpActionTraitMapping = map[string]func(*schema.RsvpActionTrait, []string){}
var customRsvpActionTraitMapping = map[string]func(*schema.RsvpActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RsvpAction", func(ctx jsonld.Context) (interface{}) {
		return NewRsvpActionFromContext(ctx)
	})



	basicRsvpActionTraitMapping["http://schema.org/additionalNumberOfGuests"] = func(x *schema.RsvpActionTrait, s []string) {
		var err error
		x.AdditionalNumberOfGuests, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	customRsvpActionTraitMapping["http://schema.org/comment"] = func(x *schema.RsvpActionTrait, ctx jsonld.Context, s string){
		var y schema.Comment
		if strings.HasPrefix(s, "_:") {
			y = NewCommentFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewComment()
			y.Id = s
		}

		x.Comment = &y

		return
	}

	customRsvpActionTraitMapping["http://schema.org/rsvpResponse"] = func(x *schema.RsvpActionTrait, ctx jsonld.Context, s string){
		var y schema.RsvpResponseType
		if strings.HasPrefix(s, "_:") {
			y = NewRsvpResponseTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRsvpResponseType()
			y.Id = s
		}

		x.RsvpResponse = &y

		return
	}
}

func NewRsvpActionFromContext(ctx jsonld.Context) (x schema.RsvpAction) {
	x.Type = "http://schema.org/RsvpAction"
	MapBasicToInformActionTrait(ctx, &x.InformActionTrait)
	MapCustomToInformActionTrait(ctx, &x.InformActionTrait)

	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRsvpActionTrait(ctx, &x.RsvpActionTrait)
	MapCustomToRsvpActionTrait(ctx, &x.RsvpActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRsvpActionTrait(ctx jsonld.Context, x *schema.RsvpActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRsvpActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRsvpActionTrait(ctx jsonld.Context, x *schema.RsvpActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRsvpActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}