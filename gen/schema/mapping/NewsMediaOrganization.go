package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNewsMediaOrganization strings.Replacer
var strconvNewsMediaOrganization strconv.NumError

var basicNewsMediaOrganizationTraitMapping = map[string]func(*schema.NewsMediaOrganizationTrait, []string){}
var customNewsMediaOrganizationTraitMapping = map[string]func(*schema.NewsMediaOrganizationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NewsMediaOrganization", func(ctx jsonld.Context) (interface{}) {
		return NewNewsMediaOrganizationFromContext(ctx)
	})













	customNewsMediaOrganizationTraitMapping["http://schema.org/actionableFeedbackPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/correctionsPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/diversityPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/diversityStaffingReport"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/ethicsPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/ownershipFundingInfo"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/unnamedSourcesPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
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

	customNewsMediaOrganizationTraitMapping["http://schema.org/missionCoveragePrioritiesPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MissionCoveragePrioritiesPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MissionCoveragePrioritiesPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MissionCoveragePrioritiesPolicy = s
		}
	}

	customNewsMediaOrganizationTraitMapping["http://schema.org/verificationFactCheckingPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.VerificationFactCheckingPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.VerificationFactCheckingPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.VerificationFactCheckingPolicy = s
		}
	}

	customNewsMediaOrganizationTraitMapping["http://schema.org/noBylinesPolicy"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NoBylinesPolicy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NoBylinesPolicy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NoBylinesPolicy = s
		}
	}

	customNewsMediaOrganizationTraitMapping["http://schema.org/masthead"] = func(x *schema.NewsMediaOrganizationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Masthead, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Masthead = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Masthead = s
		}
	}
}

func NewNewsMediaOrganizationFromContext(ctx jsonld.Context) (x schema.NewsMediaOrganization) {
	x.Type = "http://schema.org/NewsMediaOrganization"
	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToNewsMediaOrganizationTrait(ctx, &x.NewsMediaOrganizationTrait)
	MapCustomToNewsMediaOrganizationTrait(ctx, &x.NewsMediaOrganizationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNewsMediaOrganizationTrait(ctx jsonld.Context, x *schema.NewsMediaOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNewsMediaOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNewsMediaOrganizationTrait(ctx jsonld.Context, x *schema.NewsMediaOrganizationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNewsMediaOrganizationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}