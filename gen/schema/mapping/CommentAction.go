package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCommentAction strings.Replacer
var strconvCommentAction strconv.NumError

var basicCommentActionTraitMapping = map[string]func(*schema.CommentActionTrait, []string){}
var customCommentActionTraitMapping = map[string]func(*schema.CommentActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CommentAction", func(ctx jsonld.Context) (interface{}) {
		return NewCommentActionFromContext(ctx)
	})



	customCommentActionTraitMapping["http://schema.org/resultComment"] = func(x *schema.CommentActionTrait, ctx jsonld.Context, s string){
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

func NewCommentActionFromContext(ctx jsonld.Context) (x schema.CommentAction) {
	x.Type = "http://schema.org/CommentAction"
	MapBasicToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)
	MapCustomToCommunicateActionTrait(ctx, &x.CommunicateActionTrait)

	MapBasicToInteractActionTrait(ctx, &x.InteractActionTrait)
	MapCustomToInteractActionTrait(ctx, &x.InteractActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCommentActionTrait(ctx, &x.CommentActionTrait)
	MapCustomToCommentActionTrait(ctx, &x.CommentActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCommentActionTrait(ctx jsonld.Context, x *schema.CommentActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCommentActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCommentActionTrait(ctx jsonld.Context, x *schema.CommentActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCommentActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}