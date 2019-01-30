package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLegislation strings.Replacer
var strconvLegislation strconv.NumError

var basicLegislationTraitMapping = map[string]func(*schema.LegislationTrait, []string){}
var customLegislationTraitMapping = map[string]func(*schema.LegislationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Legislation", func(ctx jsonld.Context) (interface{}) {
		return NewLegislationFromContext(ctx)
	})














	customLegislationTraitMapping["http://schema.org/legislationIdentifier"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegislationIdentifier, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegislationIdentifier = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegislationIdentifier = s
		}
	}

	customLegislationTraitMapping["http://schema.org/legislationJurisdiction"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegislationJurisdiction, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegislationJurisdiction = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegislationJurisdiction = s
		}
	}

	customLegislationTraitMapping["http://schema.org/legislationType"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegislationType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegislationType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegislationType = s
		}
	}

	customLegislationTraitMapping["http://schema.org/legislationApplies"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.Legislation
		if strings.HasPrefix(s, "_:") {
			y = NewLegislationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLegislation()
			y.Id = s
		}

		x.LegislationApplies = &y

		return
	}

	customLegislationTraitMapping["http://schema.org/legislationChanges"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.Legislation
		if strings.HasPrefix(s, "_:") {
			y = NewLegislationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLegislation()
			y.Id = s
		}

		x.LegislationChanges = &y

		return
	}

	customLegislationTraitMapping["http://schema.org/legislationConsolidates"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.Legislation
		if strings.HasPrefix(s, "_:") {
			y = NewLegislationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLegislation()
			y.Id = s
		}

		x.LegislationConsolidates = &y

		return
	}

	customLegislationTraitMapping["http://schema.org/legislationDate"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.LegislationDate = &y

		return
	}

	customLegislationTraitMapping["http://schema.org/legislationDateVersion"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.LegislationDateVersion = &y

		return
	}

	customLegislationTraitMapping["http://schema.org/legislationLegalForce"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.LegalForceStatus
		if strings.HasPrefix(s, "_:") {
			y = NewLegalForceStatusFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLegalForceStatus()
			y.Id = s
		}

		x.LegislationLegalForce = &y

		return
	}

	customLegislationTraitMapping["http://schema.org/legislationPassedBy"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegislationPassedBy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegislationPassedBy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegislationPassedBy = s
		}
	}

	customLegislationTraitMapping["http://schema.org/legislationResponsible"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegislationResponsible, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegislationResponsible = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegislationResponsible = s
		}
	}

	customLegislationTraitMapping["http://schema.org/legislationTransposes"] = func(x *schema.LegislationTrait, ctx jsonld.Context, s string){
		var y schema.Legislation
		if strings.HasPrefix(s, "_:") {
			y = NewLegislationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewLegislation()
			y.Id = s
		}

		x.LegislationTransposes = &y

		return
	}
}

func NewLegislationFromContext(ctx jsonld.Context) (x schema.Legislation) {
	x.Type = "http://schema.org/Legislation"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLegislationTrait(ctx, &x.LegislationTrait)
	MapCustomToLegislationTrait(ctx, &x.LegislationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLegislationTrait(ctx jsonld.Context, x *schema.LegislationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLegislationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLegislationTrait(ctx jsonld.Context, x *schema.LegislationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLegislationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}