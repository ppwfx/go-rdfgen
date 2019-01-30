package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPermit strings.Replacer
var strconvPermit strconv.NumError

var basicPermitTraitMapping = map[string]func(*schema.PermitTrait, []string){}
var customPermitTraitMapping = map[string]func(*schema.PermitTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Permit", func(ctx jsonld.Context) (interface{}) {
		return NewPermitFromContext(ctx)
	})









	customPermitTraitMapping["http://schema.org/validFrom"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ValidFrom = &y

		return
	}

	customPermitTraitMapping["http://schema.org/issuedBy"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.IssuedBy = &y

		return
	}

	customPermitTraitMapping["http://schema.org/issuedThrough"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.Service
		if strings.HasPrefix(s, "_:") {
			y = NewServiceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewService()
			y.Id = s
		}

		x.IssuedThrough = &y

		return
	}

	customPermitTraitMapping["http://schema.org/permitAudience"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.Audience
		if strings.HasPrefix(s, "_:") {
			y = NewAudienceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudience()
			y.Id = s
		}

		x.PermitAudience = &y

		return
	}

	customPermitTraitMapping["http://schema.org/validFor"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.ValidFor = &y

		return
	}

	customPermitTraitMapping["http://schema.org/validIn"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.AdministrativeArea
		if strings.HasPrefix(s, "_:") {
			y = NewAdministrativeAreaFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAdministrativeArea()
			y.Id = s
		}

		x.ValidIn = &y

		return
	}

	customPermitTraitMapping["http://schema.org/validUntil"] = func(x *schema.PermitTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.ValidUntil = &y

		return
	}
}

func NewPermitFromContext(ctx jsonld.Context) (x schema.Permit) {
	x.Type = "http://schema.org/Permit"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPermitTrait(ctx, &x.PermitTrait)
	MapCustomToPermitTrait(ctx, &x.PermitTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPermitTrait(ctx jsonld.Context, x *schema.PermitTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPermitTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPermitTrait(ctx jsonld.Context, x *schema.PermitTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPermitTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}