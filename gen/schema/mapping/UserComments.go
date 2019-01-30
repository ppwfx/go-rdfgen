package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsUserComments strings.Replacer
var strconvUserComments strconv.NumError

var basicUserCommentsTraitMapping = map[string]func(*schema.UserCommentsTrait, []string){}
var customUserCommentsTraitMapping = map[string]func(*schema.UserCommentsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/UserComments", func(ctx jsonld.Context) (interface{}) {
		return NewUserCommentsFromContext(ctx)
	})



	basicUserCommentsTraitMapping["http://schema.org/commentText"] = func(x *schema.UserCommentsTrait, s []string) {
		x.CommentText = s[0]
	}




	basicUserCommentsTraitMapping["http://schema.org/replyToUrl"] = func(x *schema.UserCommentsTrait, s []string) {
		x.ReplyToUrl = s[0]
	}


	customUserCommentsTraitMapping["http://schema.org/creator"] = func(x *schema.UserCommentsTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Creator, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Creator = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Creator = s
		}
	}

	customUserCommentsTraitMapping["http://schema.org/commentTime"] = func(x *schema.UserCommentsTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CommentTime, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CommentTime = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CommentTime = s
		}
	}

	customUserCommentsTraitMapping["http://schema.org/discusses"] = func(x *schema.UserCommentsTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.Discusses = &y

		return
	}
}

func NewUserCommentsFromContext(ctx jsonld.Context) (x schema.UserComments) {
	x.Type = "http://schema.org/UserComments"
	MapBasicToUserInteractionTrait(ctx, &x.UserInteractionTrait)
	MapCustomToUserInteractionTrait(ctx, &x.UserInteractionTrait)

	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToUserCommentsTrait(ctx, &x.UserCommentsTrait)
	MapCustomToUserCommentsTrait(ctx, &x.UserCommentsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToUserCommentsTrait(ctx jsonld.Context, x *schema.UserCommentsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicUserCommentsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToUserCommentsTrait(ctx jsonld.Context, x *schema.UserCommentsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customUserCommentsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}