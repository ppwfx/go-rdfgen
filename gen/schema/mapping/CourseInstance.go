package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCourseInstance strings.Replacer
var strconvCourseInstance strconv.NumError

var basicCourseInstanceTraitMapping = map[string]func(*schema.CourseInstanceTrait, []string){}
var customCourseInstanceTraitMapping = map[string]func(*schema.CourseInstanceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CourseInstance", func(ctx jsonld.Context) (interface{}) {
		return NewCourseInstanceFromContext(ctx)
	})




	customCourseInstanceTraitMapping["http://schema.org/courseMode"] = func(x *schema.CourseInstanceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CourseMode, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CourseMode = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CourseMode = s
		}
	}

	customCourseInstanceTraitMapping["http://schema.org/instructor"] = func(x *schema.CourseInstanceTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Instructor = &y

		return
	}
}

func NewCourseInstanceFromContext(ctx jsonld.Context) (x schema.CourseInstance) {
	x.Type = "http://schema.org/CourseInstance"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCourseInstanceTrait(ctx, &x.CourseInstanceTrait)
	MapCustomToCourseInstanceTrait(ctx, &x.CourseInstanceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCourseInstanceTrait(ctx jsonld.Context, x *schema.CourseInstanceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCourseInstanceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCourseInstanceTrait(ctx jsonld.Context, x *schema.CourseInstanceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCourseInstanceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}