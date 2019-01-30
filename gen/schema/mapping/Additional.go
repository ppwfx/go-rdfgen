package mapping

import (
	"github.com/21stio/go-rdfgen/pkg/jsonld"
	"github.com/21stio/go-rdfgen/gen/schema"
	"strings"
)

func MapToAdditionalTrait(ctx jsonld.Context, x *schema.AdditionalTrait) () {
	f := func(ctx jsonld.Context, s string) (additional interface{}) {
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Decoded[s]
			if ok {
				additional = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

			_, ok = ctx.Subjects[s]
			if ok {
				var err error
				additional, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}
		} else {
			additional = s
		}

		return additional
	}

	x.Additional = map[string]interface{}{}
	for k, v := range ctx.Current {
		x.Additional[k] = f(ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}
