package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQuestion strings.Replacer
var strconvQuestion strconv.NumError

var basicQuestionTraitMapping = map[string]func(*schema.QuestionTrait, []string){}
var customQuestionTraitMapping = map[string]func(*schema.QuestionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Question", func(ctx jsonld.Context) (interface{}) {
		return NewQuestionFromContext(ctx)
	})







	customQuestionTraitMapping["http://schema.org/downvoteCount"] = func(x *schema.QuestionTrait, ctx jsonld.Context, s string){
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

	customQuestionTraitMapping["http://schema.org/upvoteCount"] = func(x *schema.QuestionTrait, ctx jsonld.Context, s string){
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

	customQuestionTraitMapping["http://schema.org/acceptedAnswer"] = func(x *schema.QuestionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AcceptedAnswer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AcceptedAnswer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AcceptedAnswer = s
		}
	}

	customQuestionTraitMapping["http://schema.org/answerCount"] = func(x *schema.QuestionTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.AnswerCount = &y

		return
	}

	customQuestionTraitMapping["http://schema.org/suggestedAnswer"] = func(x *schema.QuestionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SuggestedAnswer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SuggestedAnswer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SuggestedAnswer = s
		}
	}
}

func NewQuestionFromContext(ctx jsonld.Context) (x schema.Question) {
	x.Type = "http://schema.org/Question"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQuestionTrait(ctx, &x.QuestionTrait)
	MapCustomToQuestionTrait(ctx, &x.QuestionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQuestionTrait(ctx jsonld.Context, x *schema.QuestionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQuestionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQuestionTrait(ctx jsonld.Context, x *schema.QuestionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQuestionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}