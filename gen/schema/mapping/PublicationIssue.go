package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPublicationIssue strings.Replacer
var strconvPublicationIssue strconv.NumError

var basicPublicationIssueTraitMapping = map[string]func(*schema.PublicationIssueTrait, []string){}
var customPublicationIssueTraitMapping = map[string]func(*schema.PublicationIssueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PublicationIssue", func(ctx jsonld.Context) (interface{}) {
		return NewPublicationIssueFromContext(ctx)
	})




	basicPublicationIssueTraitMapping["http://schema.org/pagination"] = func(x *schema.PublicationIssueTrait, s []string) {
		x.Pagination = s[0]
	}



	customPublicationIssueTraitMapping["http://schema.org/pageEnd"] = func(x *schema.PublicationIssueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PageEnd, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PageEnd = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PageEnd = s
		}
	}

	customPublicationIssueTraitMapping["http://schema.org/pageStart"] = func(x *schema.PublicationIssueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PageStart, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PageStart = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PageStart = s
		}
	}

	customPublicationIssueTraitMapping["http://schema.org/issueNumber"] = func(x *schema.PublicationIssueTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IssueNumber, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IssueNumber = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IssueNumber = s
		}
	}
}

func NewPublicationIssueFromContext(ctx jsonld.Context) (x schema.PublicationIssue) {
	x.Type = "http://schema.org/PublicationIssue"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPublicationIssueTrait(ctx, &x.PublicationIssueTrait)
	MapCustomToPublicationIssueTrait(ctx, &x.PublicationIssueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPublicationIssueTrait(ctx jsonld.Context, x *schema.PublicationIssueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPublicationIssueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPublicationIssueTrait(ctx jsonld.Context, x *schema.PublicationIssueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPublicationIssueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}