package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCourse strings.Replacer
var strconvCourse strconv.NumError

var basicCourseTraitMapping = map[string]func(*schema.CourseTrait, []string){}
var customCourseTraitMapping = map[string]func(*schema.CourseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Course", func(ctx jsonld.Context) (interface{}) {
		return NewCourseFromContext(ctx)
	})


	basicCourseTraitMapping["http://schema.org/courseCode"] = func(x *schema.CourseTrait, s []string) {
		x.CourseCode = s[0]
	}





	customCourseTraitMapping["http://schema.org/coursePrerequisites"] = func(x *schema.CourseTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CoursePrerequisites, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CoursePrerequisites = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CoursePrerequisites = s
		}
	}

	customCourseTraitMapping["http://schema.org/educationalCredentialAwarded"] = func(x *schema.CourseTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EducationalCredentialAwarded, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EducationalCredentialAwarded = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EducationalCredentialAwarded = s
		}
	}

	customCourseTraitMapping["http://schema.org/hasCourseInstance"] = func(x *schema.CourseTrait, ctx jsonld.Context, s string){
		var y schema.CourseInstance
		if strings.HasPrefix(s, "_:") {
			y = NewCourseInstanceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCourseInstance()
			y.Id = s
		}

		x.HasCourseInstance = &y

		return
	}
}

func NewCourseFromContext(ctx jsonld.Context) (x schema.Course) {
	x.Type = "http://schema.org/Course"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCourseTrait(ctx, &x.CourseTrait)
	MapCustomToCourseTrait(ctx, &x.CourseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCourseTrait(ctx jsonld.Context, x *schema.CourseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCourseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCourseTrait(ctx jsonld.Context, x *schema.CourseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCourseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}