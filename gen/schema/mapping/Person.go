package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPerson strings.Replacer
var strconvPerson strconv.NumError

var basicPersonTraitMapping = map[string]func(*schema.PersonTrait, []string){}
var customPersonTraitMapping = map[string]func(*schema.PersonTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Person", func(ctx jsonld.Context) (interface{}) {
		return NewPersonFromContext(ctx)
	})


	basicPersonTraitMapping["http://schema.org/award"] = func(x *schema.PersonTrait, s []string) {
		x.Award = s[0]
	}


	basicPersonTraitMapping["http://schema.org/awards"] = func(x *schema.PersonTrait, s []string) {
		x.Awards = s[0]
	}







	basicPersonTraitMapping["http://schema.org/email"] = func(x *schema.PersonTrait, s []string) {
		x.Email = s[0]
	}


	basicPersonTraitMapping["http://schema.org/faxNumber"] = func(x *schema.PersonTrait, s []string) {
		x.FaxNumber = s[0]
	}


	basicPersonTraitMapping["http://schema.org/telephone"] = func(x *schema.PersonTrait, s []string) {
		x.Telephone = s[0]
	}


	basicPersonTraitMapping["http://schema.org/globalLocationNumber"] = func(x *schema.PersonTrait, s []string) {
		x.GlobalLocationNumber = s[0]
	}


	basicPersonTraitMapping["http://schema.org/isicV4"] = func(x *schema.PersonTrait, s []string) {
		x.IsicV4 = s[0]
	}


	basicPersonTraitMapping["http://schema.org/additionalName"] = func(x *schema.PersonTrait, s []string) {
		x.AdditionalName = s[0]
	}














	basicPersonTraitMapping["http://schema.org/duns"] = func(x *schema.PersonTrait, s []string) {
		x.Duns = s[0]
	}


	basicPersonTraitMapping["http://schema.org/familyName"] = func(x *schema.PersonTrait, s []string) {
		x.FamilyName = s[0]
	}




	basicPersonTraitMapping["http://schema.org/givenName"] = func(x *schema.PersonTrait, s []string) {
		x.GivenName = s[0]
	}






	basicPersonTraitMapping["http://schema.org/honorificPrefix"] = func(x *schema.PersonTrait, s []string) {
		x.HonorificPrefix = s[0]
	}


	basicPersonTraitMapping["http://schema.org/honorificSuffix"] = func(x *schema.PersonTrait, s []string) {
		x.HonorificSuffix = s[0]
	}


	basicPersonTraitMapping["http://schema.org/jobTitle"] = func(x *schema.PersonTrait, s []string) {
		x.JobTitle = s[0]
	}







	basicPersonTraitMapping["http://schema.org/naics"] = func(x *schema.PersonTrait, s []string) {
		x.Naics = s[0]
	}













	basicPersonTraitMapping["http://schema.org/taxID"] = func(x *schema.PersonTrait, s []string) {
		x.TaxID = s[0]
	}


	basicPersonTraitMapping["http://schema.org/vatID"] = func(x *schema.PersonTrait, s []string) {
		x.VatID = s[0]
	}





	customPersonTraitMapping["http://schema.org/funder"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Funder, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Funder = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Funder = s
		}
	}

	customPersonTraitMapping["http://schema.org/height"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Height, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Height = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Height = s
		}
	}

	customPersonTraitMapping["http://schema.org/publishingPrinciples"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PublishingPrinciples, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PublishingPrinciples = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PublishingPrinciples = s
		}
	}

	customPersonTraitMapping["http://schema.org/sponsor"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Sponsor, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Sponsor = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Sponsor = s
		}
	}

	customPersonTraitMapping["http://schema.org/address"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Address, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Address = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Address = s
		}
	}

	customPersonTraitMapping["http://schema.org/affiliation"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Affiliation = &y

		return
	}

	customPersonTraitMapping["http://schema.org/alumniOf"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AlumniOf, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AlumniOf = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AlumniOf = s
		}
	}

	customPersonTraitMapping["http://schema.org/birthDate"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.BirthDate = &y

		return
	}

	customPersonTraitMapping["http://schema.org/birthPlace"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.BirthPlace = &y

		return
	}

	customPersonTraitMapping["http://schema.org/brand"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Brand, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Brand = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Brand = s
		}
	}

	customPersonTraitMapping["http://schema.org/children"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Children = &y

		return
	}

	customPersonTraitMapping["http://schema.org/colleague"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Colleague, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Colleague = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Colleague = s
		}
	}

	customPersonTraitMapping["http://schema.org/colleagues"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Colleagues = &y

		return
	}

	customPersonTraitMapping["http://schema.org/contactPoint"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.ContactPoint
		if strings.HasPrefix(s, "_:") {
			y = NewContactPointFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewContactPoint()
			y.Id = s
		}

		x.ContactPoint = &y

		return
	}

	customPersonTraitMapping["http://schema.org/contactPoints"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.ContactPoint
		if strings.HasPrefix(s, "_:") {
			y = NewContactPointFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewContactPoint()
			y.Id = s
		}

		x.ContactPoints = &y

		return
	}

	customPersonTraitMapping["http://schema.org/deathDate"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.DeathDate = &y

		return
	}

	customPersonTraitMapping["http://schema.org/deathPlace"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.DeathPlace = &y

		return
	}

	customPersonTraitMapping["http://schema.org/follows"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Follows = &y

		return
	}

	customPersonTraitMapping["http://schema.org/gender"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Gender, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Gender = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Gender = s
		}
	}

	customPersonTraitMapping["http://schema.org/hasOccupation"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Occupation
		if strings.HasPrefix(s, "_:") {
			y = NewOccupationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOccupation()
			y.Id = s
		}

		x.HasOccupation = &y

		return
	}

	customPersonTraitMapping["http://schema.org/hasOfferCatalog"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.OfferCatalog
		if strings.HasPrefix(s, "_:") {
			y = NewOfferCatalogFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOfferCatalog()
			y.Id = s
		}

		x.HasOfferCatalog = &y

		return
	}

	customPersonTraitMapping["http://schema.org/hasPOS"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.HasPOS = &y

		return
	}

	customPersonTraitMapping["http://schema.org/homeLocation"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HomeLocation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HomeLocation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HomeLocation = s
		}
	}

	customPersonTraitMapping["http://schema.org/knows"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Knows = &y

		return
	}

	customPersonTraitMapping["http://schema.org/knowsAbout"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.KnowsAbout, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.KnowsAbout = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.KnowsAbout = s
		}
	}

	customPersonTraitMapping["http://schema.org/knowsLanguage"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.KnowsLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.KnowsLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.KnowsLanguage = s
		}
	}

	customPersonTraitMapping["http://schema.org/makesOffer"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.MakesOffer = &y

		return
	}

	customPersonTraitMapping["http://schema.org/memberOf"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MemberOf, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MemberOf = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MemberOf = s
		}
	}

	customPersonTraitMapping["http://schema.org/nationality"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Country
		if strings.HasPrefix(s, "_:") {
			y = NewCountryFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCountry()
			y.Id = s
		}

		x.Nationality = &y

		return
	}

	customPersonTraitMapping["http://schema.org/netWorth"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NetWorth, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NetWorth = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NetWorth = s
		}
	}

	customPersonTraitMapping["http://schema.org/owns"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Owns, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Owns = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Owns = s
		}
	}

	customPersonTraitMapping["http://schema.org/parent"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Parent = &y

		return
	}

	customPersonTraitMapping["http://schema.org/parents"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Parents = &y

		return
	}

	customPersonTraitMapping["http://schema.org/performerIn"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.PerformerIn = &y

		return
	}

	customPersonTraitMapping["http://schema.org/relatedTo"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.RelatedTo = &y

		return
	}

	customPersonTraitMapping["http://schema.org/seeks"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Demand
		if strings.HasPrefix(s, "_:") {
			y = NewDemandFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDemand()
			y.Id = s
		}

		x.Seeks = &y

		return
	}

	customPersonTraitMapping["http://schema.org/sibling"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Sibling = &y

		return
	}

	customPersonTraitMapping["http://schema.org/siblings"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Siblings = &y

		return
	}

	customPersonTraitMapping["http://schema.org/spouse"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Spouse = &y

		return
	}

	customPersonTraitMapping["http://schema.org/weight"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Weight = &y

		return
	}

	customPersonTraitMapping["http://schema.org/workLocation"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.WorkLocation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.WorkLocation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.WorkLocation = s
		}
	}

	customPersonTraitMapping["http://schema.org/worksFor"] = func(x *schema.PersonTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.WorksFor = &y

		return
	}
}

func NewPersonFromContext(ctx jsonld.Context) (x schema.Person) {
	x.Type = "http://schema.org/Person"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPersonTrait(ctx, &x.PersonTrait)
	MapCustomToPersonTrait(ctx, &x.PersonTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPersonTrait(ctx jsonld.Context, x *schema.PersonTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPersonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPersonTrait(ctx jsonld.Context, x *schema.PersonTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPersonTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}