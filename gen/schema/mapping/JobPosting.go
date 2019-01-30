package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsJobPosting strings.Replacer
var strconvJobPosting strconv.NumError

var basicJobPostingTraitMapping = map[string]func(*schema.JobPostingTrait, []string){}
var customJobPostingTraitMapping = map[string]func(*schema.JobPostingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/JobPosting", func(ctx jsonld.Context) (interface{}) {
		return NewJobPostingFromContext(ctx)
	})




	basicJobPostingTraitMapping["http://schema.org/benefits"] = func(x *schema.JobPostingTrait, s []string) {
		x.Benefits = s[0]
	}



	basicJobPostingTraitMapping["http://schema.org/educationRequirements"] = func(x *schema.JobPostingTrait, s []string) {
		x.EducationRequirements = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/employmentType"] = func(x *schema.JobPostingTrait, s []string) {
		x.EmploymentType = s[0]
	}



	basicJobPostingTraitMapping["http://schema.org/experienceRequirements"] = func(x *schema.JobPostingTrait, s []string) {
		x.ExperienceRequirements = s[0]
	}



	basicJobPostingTraitMapping["http://schema.org/incentiveCompensation"] = func(x *schema.JobPostingTrait, s []string) {
		x.IncentiveCompensation = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/incentives"] = func(x *schema.JobPostingTrait, s []string) {
		x.Incentives = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/industry"] = func(x *schema.JobPostingTrait, s []string) {
		x.Industry = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/jobBenefits"] = func(x *schema.JobPostingTrait, s []string) {
		x.JobBenefits = s[0]
	}



	basicJobPostingTraitMapping["http://schema.org/occupationalCategory"] = func(x *schema.JobPostingTrait, s []string) {
		x.OccupationalCategory = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/qualifications"] = func(x *schema.JobPostingTrait, s []string) {
		x.Qualifications = s[0]
	}



	basicJobPostingTraitMapping["http://schema.org/responsibilities"] = func(x *schema.JobPostingTrait, s []string) {
		x.Responsibilities = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/salaryCurrency"] = func(x *schema.JobPostingTrait, s []string) {
		x.SalaryCurrency = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/skills"] = func(x *schema.JobPostingTrait, s []string) {
		x.Skills = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/specialCommitments"] = func(x *schema.JobPostingTrait, s []string) {
		x.SpecialCommitments = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/title"] = func(x *schema.JobPostingTrait, s []string) {
		x.Title = s[0]
	}


	basicJobPostingTraitMapping["http://schema.org/workHours"] = func(x *schema.JobPostingTrait, s []string) {
		x.WorkHours = s[0]
	}


	customJobPostingTraitMapping["http://schema.org/validThrough"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ValidThrough = &y

		return
	}

	customJobPostingTraitMapping["http://schema.org/baseSalary"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BaseSalary, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BaseSalary = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BaseSalary = s
		}
	}

	customJobPostingTraitMapping["http://schema.org/datePosted"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.DatePosted = &y

		return
	}

	customJobPostingTraitMapping["http://schema.org/estimatedSalary"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EstimatedSalary, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EstimatedSalary = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EstimatedSalary = s
		}
	}

	customJobPostingTraitMapping["http://schema.org/hiringOrganization"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.HiringOrganization = &y

		return
	}

	customJobPostingTraitMapping["http://schema.org/jobLocation"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.JobLocation = &y

		return
	}

	customJobPostingTraitMapping["http://schema.org/relevantOccupation"] = func(x *schema.JobPostingTrait, ctx jsonld.Context, s string){
		var y schema.Occupation
		if strings.HasPrefix(s, "_:") {
			y = NewOccupationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOccupation()
			y.Id = s
		}

		x.RelevantOccupation = &y

		return
	}
}

func NewJobPostingFromContext(ctx jsonld.Context) (x schema.JobPosting) {
	x.Type = "http://schema.org/JobPosting"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToJobPostingTrait(ctx, &x.JobPostingTrait)
	MapCustomToJobPostingTrait(ctx, &x.JobPostingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToJobPostingTrait(ctx jsonld.Context, x *schema.JobPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicJobPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToJobPostingTrait(ctx jsonld.Context, x *schema.JobPostingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customJobPostingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}