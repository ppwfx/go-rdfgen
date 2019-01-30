package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrganization strings.Replacer
var strconvOrganization strconv.NumError

var basicOrganizationTraitMapping = map[string]func(*schema.OrganizationTrait, []string){}
var customOrganizationTraitMapping = map[string]func(*schema.OrganizationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Organization", func(ctx jsonld.Context) (interface{}) {
		return NewOrganizationFromContext(ctx)
	})



	basicOrganizationTraitMapping["http://schema.org/award"] = func(x *schema.OrganizationTrait, s []string) {
		x.Award = s[0]
	}


	basicOrganizationTraitMapping["http://schema.org/awards"] = func(x *schema.OrganizationTrait, s []string) {
		x.Awards = s[0]
	}










	basicOrganizationTraitMapping["http://schema.org/email"] = func(x *schema.OrganizationTrait, s []string) {
		x.Email = s[0]
	}


	basicOrganizationTraitMapping["http://schema.org/faxNumber"] = func(x *schema.OrganizationTrait, s []string) {
		x.FaxNumber = s[0]
	}



	basicOrganizationTraitMapping["http://schema.org/telephone"] = func(x *schema.OrganizationTrait, s []string) {
		x.Telephone = s[0]
	}




	basicOrganizationTraitMapping["http://schema.org/globalLocationNumber"] = func(x *schema.OrganizationTrait, s []string) {
		x.GlobalLocationNumber = s[0]
	}


	basicOrganizationTraitMapping["http://schema.org/isicV4"] = func(x *schema.OrganizationTrait, s []string) {
		x.IsicV4 = s[0]
	}






	basicOrganizationTraitMapping["http://schema.org/duns"] = func(x *schema.OrganizationTrait, s []string) {
		x.Duns = s[0]
	}








	basicOrganizationTraitMapping["http://schema.org/naics"] = func(x *schema.OrganizationTrait, s []string) {
		x.Naics = s[0]
	}




	basicOrganizationTraitMapping["http://schema.org/taxID"] = func(x *schema.OrganizationTrait, s []string) {
		x.TaxID = s[0]
	}


	basicOrganizationTraitMapping["http://schema.org/vatID"] = func(x *schema.OrganizationTrait, s []string) {
		x.VatID = s[0]
	}
















	basicOrganizationTraitMapping["http://schema.org/legalName"] = func(x *schema.OrganizationTrait, s []string) {
		x.LegalName = s[0]
	}


	basicOrganizationTraitMapping["http://schema.org/leiCode"] = func(x *schema.OrganizationTrait, s []string) {
		x.LeiCode = s[0]
	}









	customOrganizationTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.AggregateRating
		if strings.HasPrefix(s, "_:") {
			y = NewAggregateRatingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAggregateRating()
			y.Id = s
		}

		x.AggregateRating = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/funder"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/publishingPrinciples"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/review"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Review = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/reviews"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Review
		if strings.HasPrefix(s, "_:") {
			y = NewReviewFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewReview()
			y.Id = s
		}

		x.Reviews = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/sponsor"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/location"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Location, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Location = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Location = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/areaServed"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AreaServed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AreaServed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AreaServed = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/address"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/serviceArea"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ServiceArea, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ServiceArea = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ServiceArea = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/event"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.Event = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/events"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.Events = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/logo"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Logo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Logo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Logo = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/brand"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/contactPoint"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/contactPoints"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/hasOfferCatalog"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/hasPOS"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/knowsAbout"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/knowsLanguage"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/makesOffer"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/memberOf"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/owns"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/seeks"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
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

	customOrganizationTraitMapping["http://schema.org/actionableFeedbackPolicy"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ActionableFeedbackPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ActionableFeedbackPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ActionableFeedbackPolicy = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/alumni"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Alumni = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/correctionsPolicy"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CorrectionsPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CorrectionsPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CorrectionsPolicy = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/department"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Department = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/dissolutionDate"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.DissolutionDate = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/diversityPolicy"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DiversityPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DiversityPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DiversityPolicy = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/diversityStaffingReport"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DiversityStaffingReport, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DiversityStaffingReport = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DiversityStaffingReport = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/employee"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Employee = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/employees"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Employees = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/ethicsPolicy"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EthicsPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EthicsPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EthicsPolicy = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/founder"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Founder = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/founders"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Founders = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/foundingDate"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.FoundingDate = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/foundingLocation"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.FoundingLocation = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/member"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Member, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Member = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Member = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/members"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Members, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Members = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Members = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/numberOfEmployees"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.NumberOfEmployees = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/ownershipFundingInfo"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.OwnershipFundingInfo, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.OwnershipFundingInfo = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.OwnershipFundingInfo = s
		}
	}

	customOrganizationTraitMapping["http://schema.org/parentOrganization"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.ParentOrganization = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/subOrganization"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.SubOrganization = &y

		return
	}

	customOrganizationTraitMapping["http://schema.org/unnamedSourcesPolicy"] = func(x *schema.OrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UnnamedSourcesPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UnnamedSourcesPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UnnamedSourcesPolicy = s
		}
	}
}

func NewOrganizationFromContext(ctx jsonld.Context) (x schema.Organization) {
	x.Type = "http://schema.org/Organization"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrganizationTrait(ctx jsonld.Context, x *schema.OrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrganizationTrait(ctx jsonld.Context, x *schema.OrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}