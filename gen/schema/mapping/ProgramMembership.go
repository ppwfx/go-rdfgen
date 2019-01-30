package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsProgramMembership strings.Replacer
var strconvProgramMembership strconv.NumError

var basicProgramMembershipTraitMapping = map[string]func(*schema.ProgramMembershipTrait, []string){}
var customProgramMembershipTraitMapping = map[string]func(*schema.ProgramMembershipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ProgramMembership", func(ctx jsonld.Context) (interface{}) {
		return NewProgramMembershipFromContext(ctx)
	})





	basicProgramMembershipTraitMapping["http://schema.org/membershipNumber"] = func(x *schema.ProgramMembershipTrait, s []string) {
		x.MembershipNumber = s[0]
	}


	basicProgramMembershipTraitMapping["http://schema.org/programName"] = func(x *schema.ProgramMembershipTrait, s []string) {
		x.ProgramName = s[0]
	}


	customProgramMembershipTraitMapping["http://schema.org/member"] = func(x *schema.ProgramMembershipTrait, ctx jsonld.Context, s string){
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

	customProgramMembershipTraitMapping["http://schema.org/members"] = func(x *schema.ProgramMembershipTrait, ctx jsonld.Context, s string){
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

	customProgramMembershipTraitMapping["http://schema.org/hostingOrganization"] = func(x *schema.ProgramMembershipTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.HostingOrganization = &y

		return
	}
}

func NewProgramMembershipFromContext(ctx jsonld.Context) (x schema.ProgramMembership) {
	x.Type = "http://schema.org/ProgramMembership"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToProgramMembershipTrait(ctx, &x.ProgramMembershipTrait)
	MapCustomToProgramMembershipTrait(ctx, &x.ProgramMembershipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToProgramMembershipTrait(ctx jsonld.Context, x *schema.ProgramMembershipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicProgramMembershipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToProgramMembershipTrait(ctx jsonld.Context, x *schema.ProgramMembershipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customProgramMembershipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}