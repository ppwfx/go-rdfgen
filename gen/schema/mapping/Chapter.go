package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsChapter strings.Replacer
var strconvChapter strconv.NumError

var basicChapterTraitMapping = map[string]func(*schema.ChapterTrait, []string){}
var customChapterTraitMapping = map[string]func(*schema.ChapterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Chapter", func(ctx jsonld.Context) (interface{}) {
		return NewChapterFromContext(ctx)
	})




	basicChapterTraitMapping["http://schema.org/pagination"] = func(x *schema.ChapterTrait, s []string) {
		x.Pagination = s[0]
	}


	customChapterTraitMapping["http://schema.org/pageEnd"] = func(x *schema.ChapterTrait, ctx jsonld.Context, s string){
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

	customChapterTraitMapping["http://schema.org/pageStart"] = func(x *schema.ChapterTrait, ctx jsonld.Context, s string){
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
}

func NewChapterFromContext(ctx jsonld.Context) (x schema.Chapter) {
	x.Type = "http://schema.org/Chapter"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToChapterTrait(ctx, &x.ChapterTrait)
	MapCustomToChapterTrait(ctx, &x.ChapterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToChapterTrait(ctx jsonld.Context, x *schema.ChapterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicChapterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToChapterTrait(ctx jsonld.Context, x *schema.ChapterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customChapterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}