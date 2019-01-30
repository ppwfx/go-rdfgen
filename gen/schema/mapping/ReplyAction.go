package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReplyAction strings.Replacer
var strconvReplyAction strconv.NumError

var basicReplyActionTraitMapping = map[string]func(*schema.ReplyActionTrait, []string){}
var customReplyActionTraitMapping = map[string]func(*schema.ReplyActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReplyAction", func(ctx jsonld.Context) (interface{}) {
		return NewReplyActionFromContext(ctx)
	})



	customReplyActionTraitMapping["http://schema.org/resultComment"] = func(x *schema.ReplyActionTrait, ctx jsonld.Context, s string){
		var y schema.Comment
		if strings.HasPrefix(s, "_:") {
			y = NewCommentFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewComment()
			y.Id = s
		}

		x.ResultComment = &y

		return
	}
}

func NewReplyActionFromContext(ctx jsonld.Context) (x schema.ReplyAction) {
	x.Type = "http://schema.org/ReplyAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReplyActionTrait(ctx, &x.ReplyActionTrait)
	MapCustomToReplyActionTrait(ctx, &x.ReplyActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReplyActionTrait(ctx jsonld.Context, x *schema.ReplyActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReplyActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReplyActionTrait(ctx jsonld.Context, x *schema.ReplyActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReplyActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}