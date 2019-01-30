package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBlog strings.Replacer
var strconvBlog strconv.NumError

var basicBlogTraitMapping = map[string]func(*schema.BlogTrait, []string){}
var customBlogTraitMapping = map[string]func(*schema.BlogTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Blog", func(ctx jsonld.Context) (interface{}) {
		return NewBlogFromContext(ctx)
	})


	basicBlogTraitMapping["http://schema.org/issn"] = func(x *schema.BlogTrait, s []string) {
		x.Issn = s[0]
	}




	customBlogTraitMapping["http://schema.org/blogPost"] = func(x *schema.BlogTrait, ctx jsonld.Context, s string){
		var y schema.BlogPosting
		if strings.HasPrefix(s, "_:") {
			y = NewBlogPostingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBlogPosting()
			y.Id = s
		}

		x.BlogPost = &y

		return
	}

	customBlogTraitMapping["http://schema.org/blogPosts"] = func(x *schema.BlogTrait, ctx jsonld.Context, s string){
		var y schema.BlogPosting
		if strings.HasPrefix(s, "_:") {
			y = NewBlogPostingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewBlogPosting()
			y.Id = s
		}

		x.BlogPosts = &y

		return
	}
}

func NewBlogFromContext(ctx jsonld.Context) (x schema.Blog) {
	x.Type = "http://schema.org/Blog"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBlogTrait(ctx, &x.BlogTrait)
	MapCustomToBlogTrait(ctx, &x.BlogTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBlogTrait(ctx jsonld.Context, x *schema.BlogTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBlogTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBlogTrait(ctx jsonld.Context, x *schema.BlogTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBlogTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}