package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEvent strings.Replacer
var strconvEvent strconv.NumError

var basicEventTraitMapping = map[string]func(*schema.EventTrait, []string){}
var customEventTraitMapping = map[string]func(*schema.EventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Event", func(ctx jsonld.Context) (interface{}) {
		return NewEventFromContext(ctx)
	})









	basicEventTraitMapping["http://schema.org/isAccessibleForFree"] = func(x *schema.EventTrait, s []string) {
		var err error
		x.IsAccessibleForFree, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}






	basicEventTraitMapping["http://schema.org/typicalAgeRange"] = func(x *schema.EventTrait, s []string) {
		x.TypicalAgeRange = s[0]
	}
























	customEventTraitMapping["http://schema.org/about"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.About = &y

		return
	}

	customEventTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
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

	customEventTraitMapping["http://schema.org/audience"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Audience
		if strings.HasPrefix(s, "_:") {
			y = NewAudienceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudience()
			y.Id = s
		}

		x.Audience = &y

		return
	}

	customEventTraitMapping["http://schema.org/contributor"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Contributor, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Contributor = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Contributor = s
		}
	}

	customEventTraitMapping["http://schema.org/duration"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.Duration = &y

		return
	}

	customEventTraitMapping["http://schema.org/funder"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
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

	customEventTraitMapping["http://schema.org/inLanguage"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.InLanguage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.InLanguage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.InLanguage = s
		}
	}

	customEventTraitMapping["http://schema.org/offers"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.Offers = &y

		return
	}

	customEventTraitMapping["http://schema.org/review"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
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

	customEventTraitMapping["http://schema.org/sponsor"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
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

	customEventTraitMapping["http://schema.org/translator"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Translator, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Translator = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Translator = s
		}
	}

	customEventTraitMapping["http://schema.org/actor"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Actor = &y

		return
	}

	customEventTraitMapping["http://schema.org/attendee"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Attendee, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Attendee = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Attendee = s
		}
	}

	customEventTraitMapping["http://schema.org/attendees"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Attendees, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Attendees = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Attendees = s
		}
	}

	customEventTraitMapping["http://schema.org/composer"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Composer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Composer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Composer = s
		}
	}

	customEventTraitMapping["http://schema.org/director"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Director = &y

		return
	}

	customEventTraitMapping["http://schema.org/doorTime"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DoorTime = &y

		return
	}

	customEventTraitMapping["http://schema.org/endDate"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EndDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EndDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EndDate = s
		}
	}

	customEventTraitMapping["http://schema.org/eventStatus"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.EventStatusType
		if strings.HasPrefix(s, "_:") {
			y = NewEventStatusTypeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEventStatusType()
			y.Id = s
		}

		x.EventStatus = &y

		return
	}

	customEventTraitMapping["http://schema.org/location"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
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

	customEventTraitMapping["http://schema.org/maximumAttendeeCapacity"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.MaximumAttendeeCapacity = &y

		return
	}

	customEventTraitMapping["http://schema.org/organizer"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Organizer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Organizer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Organizer = s
		}
	}

	customEventTraitMapping["http://schema.org/performer"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Performer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Performer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Performer = s
		}
	}

	customEventTraitMapping["http://schema.org/performers"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Performers, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Performers = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Performers = s
		}
	}

	customEventTraitMapping["http://schema.org/previousStartDate"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.PreviousStartDate = &y

		return
	}

	customEventTraitMapping["http://schema.org/recordedIn"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.RecordedIn = &y

		return
	}

	customEventTraitMapping["http://schema.org/remainingAttendeeCapacity"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.RemainingAttendeeCapacity = &y

		return
	}

	customEventTraitMapping["http://schema.org/startDate"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.StartDate, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.StartDate = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.StartDate = s
		}
	}

	customEventTraitMapping["http://schema.org/subEvent"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.SubEvent = &y

		return
	}

	customEventTraitMapping["http://schema.org/subEvents"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.SubEvents = &y

		return
	}

	customEventTraitMapping["http://schema.org/superEvent"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.SuperEvent = &y

		return
	}

	customEventTraitMapping["http://schema.org/workFeatured"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.WorkFeatured = &y

		return
	}

	customEventTraitMapping["http://schema.org/workPerformed"] = func(x *schema.EventTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.WorkPerformed = &y

		return
	}
}

func NewEventFromContext(ctx jsonld.Context) (x schema.Event) {
	x.Type = "http://schema.org/Event"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEventTrait(ctx jsonld.Context, x *schema.EventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEventTrait(ctx jsonld.Context, x *schema.EventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}