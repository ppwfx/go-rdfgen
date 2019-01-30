package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCorrectionComment strings.Replacer
var strconvCorrectionComment strconv.NumError

var basicCorrectionCommentTraitMapping = map[string]func(*schema.CorrectionCommentTrait, []string){}
var customCorrectionCommentTraitMapping = map[string]func(*schema.CorrectionCommentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CorrectionComment", func(ctx jsonld.Context) (interface{}) {
		return NewCorrectionCommentFromContext(ctx)
	})

}

func NewCorrectionCommentFromContext(ctx jsonld.Context) (x schema.CorrectionComment) {
	x.Type = "http://schema.org/CorrectionComment"
	MapBasicToCommentTrait(ctx, &x.CommentTrait)
	MapCustomToCommentTrait(ctx, &x.CommentTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCorrectionCommentTrait(ctx, &x.CorrectionCommentTrait)
	MapCustomToCorrectionCommentTrait(ctx, &x.CorrectionCommentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCorrectionCommentTrait(ctx jsonld.Context, x *schema.CorrectionCommentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCorrectionCommentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCorrectionCommentTrait(ctx jsonld.Context, x *schema.CorrectionCommentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCorrectionCommentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}