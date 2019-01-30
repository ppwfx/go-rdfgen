package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTicket strings.Replacer
var strconvTicket strconv.NumError

var basicTicketTraitMapping = map[string]func(*schema.TicketTrait, []string){}
var customTicketTraitMapping = map[string]func(*schema.TicketTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Ticket", func(ctx jsonld.Context) (interface{}) {
		return NewTicketFromContext(ctx)
	})



	basicTicketTraitMapping["http://schema.org/priceCurrency"] = func(x *schema.TicketTrait, s []string) {
		x.PriceCurrency = s[0]
	}





	basicTicketTraitMapping["http://schema.org/ticketNumber"] = func(x *schema.TicketTrait, s []string) {
		x.TicketNumber = s[0]
	}




	customTicketTraitMapping["http://schema.org/issuedBy"] = func(x *schema.TicketTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.IssuedBy = &y

		return
	}

	customTicketTraitMapping["http://schema.org/totalPrice"] = func(x *schema.TicketTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TotalPrice, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TotalPrice = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TotalPrice = s
		}
	}

	customTicketTraitMapping["http://schema.org/underName"] = func(x *schema.TicketTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UnderName, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UnderName = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UnderName = s
		}
	}

	customTicketTraitMapping["http://schema.org/dateIssued"] = func(x *schema.TicketTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DateIssued = &y

		return
	}

	customTicketTraitMapping["http://schema.org/ticketToken"] = func(x *schema.TicketTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TicketToken, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TicketToken = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TicketToken = s
		}
	}

	customTicketTraitMapping["http://schema.org/ticketedSeat"] = func(x *schema.TicketTrait, ctx jsonld.Context, s string){
		var y schema.Seat
		if strings.HasPrefix(s, "_:") {
			y = NewSeatFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSeat()
			y.Id = s
		}

		x.TicketedSeat = &y

		return
	}
}

func NewTicketFromContext(ctx jsonld.Context) (x schema.Ticket) {
	x.Type = "http://schema.org/Ticket"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTicketTrait(ctx, &x.TicketTrait)
	MapCustomToTicketTrait(ctx, &x.TicketTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTicketTrait(ctx jsonld.Context, x *schema.TicketTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTicketTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTicketTrait(ctx jsonld.Context, x *schema.TicketTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTicketTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}