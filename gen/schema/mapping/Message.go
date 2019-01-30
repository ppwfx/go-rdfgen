package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMessage strings.Replacer
var strconvMessage strconv.NumError

var basicMessageTraitMapping = map[string]func(*schema.MessageTrait, []string){}
var customMessageTraitMapping = map[string]func(*schema.MessageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Message", func(ctx jsonld.Context) (interface{}) {
		return NewMessageFromContext(ctx)
	})











	customMessageTraitMapping["http://schema.org/recipient"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Recipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Recipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Recipient = s
		}
	}

	customMessageTraitMapping["http://schema.org/dateReceived"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DateReceived = &y

		return
	}

	customMessageTraitMapping["http://schema.org/dateRead"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DateRead = &y

		return
	}

	customMessageTraitMapping["http://schema.org/dateSent"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DateSent = &y

		return
	}

	customMessageTraitMapping["http://schema.org/bccRecipient"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BccRecipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BccRecipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BccRecipient = s
		}
	}

	customMessageTraitMapping["http://schema.org/ccRecipient"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CcRecipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CcRecipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CcRecipient = s
		}
	}

	customMessageTraitMapping["http://schema.org/messageAttachment"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.MessageAttachment = &y

		return
	}

	customMessageTraitMapping["http://schema.org/sender"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Sender, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Sender = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Sender = s
		}
	}

	customMessageTraitMapping["http://schema.org/toRecipient"] = func(x *schema.MessageTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ToRecipient, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ToRecipient = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ToRecipient = s
		}
	}
}

func NewMessageFromContext(ctx jsonld.Context) (x schema.Message) {
	x.Type = "http://schema.org/Message"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMessageTrait(ctx, &x.MessageTrait)
	MapCustomToMessageTrait(ctx, &x.MessageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMessageTrait(ctx jsonld.Context, x *schema.MessageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMessageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMessageTrait(ctx jsonld.Context, x *schema.MessageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMessageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}