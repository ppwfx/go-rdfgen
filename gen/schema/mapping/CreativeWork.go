package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCreativeWork strings.Replacer
var strconvCreativeWork strconv.NumError

var basicCreativeWorkTraitMapping = map[string]func(*schema.CreativeWorkTrait, []string){}
var customCreativeWorkTraitMapping = map[string]func(*schema.CreativeWorkTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CreativeWork", func(ctx jsonld.Context) (interface{}) {
		return NewCreativeWorkFromContext(ctx)
	})



	basicCreativeWorkTraitMapping["http://schema.org/accessMode"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessMode = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/accessModeSufficient"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessModeSufficient = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/accessibilityAPI"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessibilityAPI = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/accessibilityControl"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessibilityControl = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/accessibilityFeature"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessibilityFeature = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/accessibilityHazard"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessibilityHazard = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/accessibilitySummary"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AccessibilitySummary = s[0]
	}




	basicCreativeWorkTraitMapping["http://schema.org/alternativeHeadline"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.AlternativeHeadline = s[0]
	}






	basicCreativeWorkTraitMapping["http://schema.org/award"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.Award = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/awards"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.Awards = s[0]
	}











	basicCreativeWorkTraitMapping["http://schema.org/copyrightYear"] = func(x *schema.CreativeWorkTrait, s []string) {
		var err error
		x.CopyrightYear, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}







	basicCreativeWorkTraitMapping["http://schema.org/discussionUrl"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.DiscussionUrl = s[0]
	}




	basicCreativeWorkTraitMapping["http://schema.org/educationalUse"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.EducationalUse = s[0]
	}











	basicCreativeWorkTraitMapping["http://schema.org/headline"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.Headline = s[0]
	}




	basicCreativeWorkTraitMapping["http://schema.org/interactivityType"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.InteractivityType = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/isAccessibleForFree"] = func(x *schema.CreativeWorkTrait, s []string) {
		var err error
		x.IsAccessibleForFree, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}




	basicCreativeWorkTraitMapping["http://schema.org/isFamilyFriendly"] = func(x *schema.CreativeWorkTrait, s []string) {
		var err error
		x.IsFamilyFriendly, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}



	basicCreativeWorkTraitMapping["http://schema.org/keywords"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.Keywords = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/learningResourceType"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.LearningResourceType = s[0]
	}



























	basicCreativeWorkTraitMapping["http://schema.org/text"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.Text = s[0]
	}


	basicCreativeWorkTraitMapping["http://schema.org/thumbnailUrl"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.ThumbnailUrl = s[0]
	}





	basicCreativeWorkTraitMapping["http://schema.org/typicalAgeRange"] = func(x *schema.CreativeWorkTrait, s []string) {
		x.TypicalAgeRange = s[0]
	}






	customCreativeWorkTraitMapping["http://schema.org/about"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/accountablePerson"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.AccountablePerson = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/aggregateRating"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/associatedMedia"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.MediaObject
		if strings.HasPrefix(s, "_:") {
			y = NewMediaObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMediaObject()
			y.Id = s
		}

		x.AssociatedMedia = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/audience"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/audio"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.AudioObject
		if strings.HasPrefix(s, "_:") {
			y = NewAudioObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAudioObject()
			y.Id = s
		}

		x.Audio = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/author"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Author, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Author = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Author = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/character"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Character = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/citation"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Citation, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Citation = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Citation = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/comment"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Comment
		if strings.HasPrefix(s, "_:") {
			y = NewCommentFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewComment()
			y.Id = s
		}

		x.Comment = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/commentCount"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Integer
		if strings.HasPrefix(s, "_:") {
			y = NewIntegerFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteger()
			y.Id = s
		}

		x.CommentCount = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/contentLocation"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.ContentLocation = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/contentRating"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ContentRating, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ContentRating = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ContentRating = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/contentReferenceTime"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ContentReferenceTime = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/contributor"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/copyrightHolder"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CopyrightHolder, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CopyrightHolder = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CopyrightHolder = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/correction"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Correction, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Correction = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Correction = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/creator"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Creator, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Creator = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Creator = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/dateCreated"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DateCreated, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DateCreated = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DateCreated = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/dateModified"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DateModified, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DateModified = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DateModified = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/datePublished"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.DatePublished = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/editor"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Editor = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/educationalAlignment"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.AlignmentObject
		if strings.HasPrefix(s, "_:") {
			y = NewAlignmentObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAlignmentObject()
			y.Id = s
		}

		x.EducationalAlignment = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/encoding"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.MediaObject
		if strings.HasPrefix(s, "_:") {
			y = NewMediaObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMediaObject()
			y.Id = s
		}

		x.Encoding = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/encodingFormat"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.EncodingFormat, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.EncodingFormat = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.EncodingFormat = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/encodings"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.MediaObject
		if strings.HasPrefix(s, "_:") {
			y = NewMediaObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMediaObject()
			y.Id = s
		}

		x.Encodings = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/exampleOfWork"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.ExampleOfWork = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/expires"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.Expires = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/fileFormat"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FileFormat, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FileFormat = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FileFormat = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/funder"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/genre"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Genre, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Genre = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Genre = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/hasPart"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HasPart, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HasPart = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HasPart = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/inLanguage"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/interactionStatistic"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.InteractionCounter
		if strings.HasPrefix(s, "_:") {
			y = NewInteractionCounterFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewInteractionCounter()
			y.Id = s
		}

		x.InteractionStatistic = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/isBasedOn"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsBasedOn, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsBasedOn = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsBasedOn = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/isBasedOnUrl"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsBasedOnUrl, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsBasedOnUrl = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsBasedOnUrl = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/isPartOf"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsPartOf, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsPartOf = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsPartOf = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/license"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.License, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.License = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.License = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/locationCreated"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.LocationCreated = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/mainEntity"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.MainEntity = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/material"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Material, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Material = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Material = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/mentions"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Thing
		if strings.HasPrefix(s, "_:") {
			y = NewThingFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewThing()
			y.Id = s
		}

		x.Mentions = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/offers"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/position"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Position, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Position = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Position = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/producer"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Producer, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Producer = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Producer = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/provider"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Provider, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Provider = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Provider = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/publication"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.PublicationEvent
		if strings.HasPrefix(s, "_:") {
			y = NewPublicationEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPublicationEvent()
			y.Id = s
		}

		x.Publication = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/publisher"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Publisher, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Publisher = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Publisher = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/publisherImprint"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.PublisherImprint = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/publishingPrinciples"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/recordedAt"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Event
		if strings.HasPrefix(s, "_:") {
			y = NewEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEvent()
			y.Id = s
		}

		x.RecordedAt = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/releasedEvent"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.PublicationEvent
		if strings.HasPrefix(s, "_:") {
			y = NewPublicationEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPublicationEvent()
			y.Id = s
		}

		x.ReleasedEvent = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/review"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/reviews"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/schemaVersion"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SchemaVersion, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SchemaVersion = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SchemaVersion = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/sdDatePublished"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.SdDatePublished = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/sdLicense"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SdLicense, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SdLicense = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SdLicense = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/sdPublisher"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SdPublisher, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SdPublisher = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SdPublisher = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/sourceOrganization"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.SourceOrganization = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/spatialCoverage"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.SpatialCoverage = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/sponsor"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/temporalCoverage"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TemporalCoverage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TemporalCoverage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TemporalCoverage = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/timeRequired"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.TimeRequired = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/translationOfWork"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.TranslationOfWork = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/translator"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
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

	customCreativeWorkTraitMapping["http://schema.org/version"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Version, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Version = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Version = s
		}
	}

	customCreativeWorkTraitMapping["http://schema.org/video"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.VideoObject
		if strings.HasPrefix(s, "_:") {
			y = NewVideoObjectFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewVideoObject()
			y.Id = s
		}

		x.Video = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/workExample"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.WorkExample = &y

		return
	}

	customCreativeWorkTraitMapping["http://schema.org/workTranslation"] = func(x *schema.CreativeWorkTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.WorkTranslation = &y

		return
	}
}

func NewCreativeWorkFromContext(ctx jsonld.Context) (x schema.CreativeWork) {
	x.Type = "http://schema.org/CreativeWork"
	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCreativeWorkTrait(ctx jsonld.Context, x *schema.CreativeWorkTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCreativeWorkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCreativeWorkTrait(ctx jsonld.Context, x *schema.CreativeWorkTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCreativeWorkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}