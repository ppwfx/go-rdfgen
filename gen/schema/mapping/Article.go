package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsArticle strings.Replacer
var strconvArticle strconv.NumError

var basicArticleTraitMapping = map[string]func(*schema.ArticleTrait, []string){}
var customArticleTraitMapping = map[string]func(*schema.ArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Article", func(ctx jsonld.Context) (interface{}) {
		return NewArticleFromContext(ctx)
	})


	basicArticleTraitMapping["http://schema.org/articleBody"] = func(x *schema.ArticleTrait, s []string) {
		x.ArticleBody = s[0]
	}


	basicArticleTraitMapping["http://schema.org/articleSection"] = func(x *schema.ArticleTrait, s []string) {
		x.ArticleSection = s[0]
	}





	basicArticleTraitMapping["http://schema.org/pagination"] = func(x *schema.ArticleTrait, s []string) {
		x.Pagination = s[0]
	}




	customArticleTraitMapping["http://schema.org/backstory"] = func(x *schema.ArticleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Backstory, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Backstory = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Backstory = s
		}
	}

	customArticleTraitMapping["http://schema.org/pageEnd"] = func(x *schema.ArticleTrait, ctx jsonld.Context, s string){
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

	customArticleTraitMapping["http://schema.org/pageStart"] = func(x *schema.ArticleTrait, ctx jsonld.Context, s string){
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

	customArticleTraitMapping["http://schema.org/speakable"] = func(x *schema.ArticleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Speakable, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Speakable = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Speakable = s
		}
	}

	customArticleTraitMapping["http://schema.org/wordCount"] = func(x *schema.ArticleTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.WordCount = &y

		return
	}
}

func NewArticleFromContext(ctx jsonld.Context) (x schema.Article) {
	x.Type = "http://schema.org/Article"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToArticleTrait(ctx jsonld.Context, x *schema.ArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToArticleTrait(ctx jsonld.Context, x *schema.ArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}