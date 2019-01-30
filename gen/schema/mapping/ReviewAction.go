package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReviewAction strings.Replacer
var strconvReviewAction strconv.NumError

var basicReviewActionTraitMapping = map[string]func(*schema.ReviewActionTrait, []string){}
var customReviewActionTraitMapping = map[string]func(*schema.ReviewActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReviewAction", func(ctx jsonld.Context) (interface{}) {
		return NewReviewActionFromContext(ctx)
	})



	customReviewActionTraitMapping["http://schema.org/resultReview"] = func(x *schema.ReviewActionTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.ResultReview = &y

		return
	}
}

func NewReviewActionFromContext(ctx jsonld.Context) (x schema.ReviewAction) {
	x.Type = "http://schema.org/ReviewAction"
	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReviewActionTrait(ctx, &x.ReviewActionTrait)
	MapCustomToReviewActionTrait(ctx, &x.ReviewActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReviewActionTrait(ctx jsonld.Context, x *schema.ReviewActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReviewActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReviewActionTrait(ctx jsonld.Context, x *schema.ReviewActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReviewActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}