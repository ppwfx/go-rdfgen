package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAction strings.Replacer
var strconvAction strconv.NumError

var basicActionTraitMapping = map[string]func(*schema.ActionTrait, []string){}
var customActionTraitMapping = map[string]func(*schema.ActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Action", func(ctx jsonld.Context) (interface{}) {
		return NewActionFromContext(ctx)
	})













	customActionTraitMapping["http://schema.org/location"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
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

	customActionTraitMapping["http://schema.org/actionStatus"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.ActionStatusType
		if strings.HasPrefix(s, "_:") {
			y = NewActionStatusTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewActionStatusType()
			y.Id = s
		}

		x.ActionStatus = &y

		return
	}

	customActionTraitMapping["http://schema.org/agent"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Agent, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Agent = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Agent = s
		}
	}

	customActionTraitMapping["http://schema.org/endTime"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.EndTime = &y

		return
	}

	customActionTraitMapping["http://schema.org/error"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Error = &y

		return
	}

	customActionTraitMapping["http://schema.org/instrument"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Instrument = &y

		return
	}

	customActionTraitMapping["http://schema.org/object"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Object = &y

		return
	}

	customActionTraitMapping["http://schema.org/participant"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Participant, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Participant = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Participant = s
		}
	}

	customActionTraitMapping["http://schema.org/result"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Result = &y

		return
	}

	customActionTraitMapping["http://schema.org/startTime"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.StartTime = &y

		return
	}

	customActionTraitMapping["http://schema.org/target"] = func(x *schema.ActionTrait, ctx jsonld.Context, s string){
		var y schema.EntryPoint
		if strings.HasPrefix(s, "_:") {
			y = NewEntryPointFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEntryPoint()
			y.Id = s
		}

		x.Target = &y

		return
	}
}

func NewActionFromContext(ctx jsonld.Context) (x schema.Action) {
	x.Type = "http://schema.org/Action"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToActionTrait(ctx jsonld.Context, x *schema.ActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToActionTrait(ctx jsonld.Context, x *schema.ActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}