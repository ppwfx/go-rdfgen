package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReport strings.Replacer
var strconvReport strconv.NumError

var basicReportTraitMapping = map[string]func(*schema.ReportTrait, []string){}
var customReportTraitMapping = map[string]func(*schema.ReportTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Report", func(ctx jsonld.Context) (interface{}) {
		return NewReportFromContext(ctx)
	})


	basicReportTraitMapping["http://schema.org/reportNumber"] = func(x *schema.ReportTrait, s []string) {
		x.ReportNumber = s[0]
	}

}

func NewReportFromContext(ctx jsonld.Context) (x schema.Report) {
	x.Type = "http://schema.org/Report"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReportTrait(ctx, &x.ReportTrait)
	MapCustomToReportTrait(ctx, &x.ReportTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReportTrait(ctx jsonld.Context, x *schema.ReportTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReportTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReportTrait(ctx jsonld.Context, x *schema.ReportTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReportTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}