package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsComment strings.Replacer
var strconvComment strconv.NumError

var basicCommentTraitMapping = map[string]func(*schema.CommentTrait, []string){}
var customCommentTraitMapping = map[string]func(*schema.CommentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Comment", func(ctx jsonld.Context) (interface{}) {
		return NewCommentFromContext(ctx)
	})





	customCommentTraitMapping["http://schema.org/downvoteCount"] = func(x *schema.CommentTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.DownvoteCount = &y

		return
	}

	customCommentTraitMapping["http://schema.org/parentItem"] = func(x *schema.CommentTrait, ctx jsonld.Context, s string){
		var y schema.Question
		if strings.HasPrefix(s, "_:") {
			y = NewQuestionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuestion()
			y.Id = s
		}

		x.ParentItem = &y

		return
	}

	customCommentTraitMapping["http://schema.org/upvoteCount"] = func(x *schema.CommentTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.UpvoteCount = &y

		return
	}
}

func NewCommentFromContext(ctx jsonld.Context) (x schema.Comment) {
	x.Type = "http://schema.org/Comment"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCommentTrait(ctx, &x.CommentTrait)
	MapCustomToCommentTrait(ctx, &x.CommentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCommentTrait(ctx jsonld.Context, x *schema.CommentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCommentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCommentTrait(ctx jsonld.Context, x *schema.CommentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCommentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}