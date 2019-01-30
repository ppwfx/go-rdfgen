package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDigitalDocumentPermission strings.Replacer
var strconvDigitalDocumentPermission strconv.NumError

var basicDigitalDocumentPermissionTraitMapping = map[string]func(*schema.DigitalDocumentPermissionTrait, []string){}
var customDigitalDocumentPermissionTraitMapping = map[string]func(*schema.DigitalDocumentPermissionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DigitalDocumentPermission", func(ctx jsonld.Context) (interface{}) {
		return NewDigitalDocumentPermissionFromContext(ctx)
	})




	customDigitalDocumentPermissionTraitMapping["http://schema.org/grantee"] = func(x *schema.DigitalDocumentPermissionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Grantee, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Grantee = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Grantee = s
		}
	}

	customDigitalDocumentPermissionTraitMapping["http://schema.org/permissionType"] = func(x *schema.DigitalDocumentPermissionTrait, ctx jsonld.Context, s string){
		var y schema.DigitalDocumentPermissionType
		if strings.HasPrefix(s, "_:") {
			y = NewDigitalDocumentPermissionTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDigitalDocumentPermissionType()
			y.Id = s
		}

		x.PermissionType = &y

		return
	}
}

func NewDigitalDocumentPermissionFromContext(ctx jsonld.Context) (x schema.DigitalDocumentPermission) {
	x.Type = "http://schema.org/DigitalDocumentPermission"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDigitalDocumentPermissionTrait(ctx, &x.DigitalDocumentPermissionTrait)
	MapCustomToDigitalDocumentPermissionTrait(ctx, &x.DigitalDocumentPermissionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDigitalDocumentPermissionTrait(ctx jsonld.Context, x *schema.DigitalDocumentPermissionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDigitalDocumentPermissionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDigitalDocumentPermissionTrait(ctx jsonld.Context, x *schema.DigitalDocumentPermissionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDigitalDocumentPermissionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}