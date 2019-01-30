package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRentAction strings.Replacer
var strconvRentAction strconv.NumError

var basicRentActionTraitMapping = map[string]func(*schema.RentActionTrait, []string){}
var customRentActionTraitMapping = map[string]func(*schema.RentActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RentAction", func(ctx jsonld.Context) (interface{}) {
		return NewRentActionFromContext(ctx)
	})




	customRentActionTraitMapping["http://schema.org/landlord"] = func(x *schema.RentActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Landlord, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Landlord = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Landlord = s
		}
	}

	customRentActionTraitMapping["http://schema.org/realEstateAgent"] = func(x *schema.RentActionTrait, ctx jsonld.Context, s string){
		var y schema.RealEstateAgent
		if strings.HasPrefix(s, "_:") {
			y = NewRealEstateAgentFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRealEstateAgent()
			y.Id = s
		}

		x.RealEstateAgent = &y

		return
	}
}

func NewRentActionFromContext(ctx jsonld.Context) (x schema.RentAction) {
	x.Type = "http://schema.org/RentAction"
	MapBasicToTradeActionTrait(ctx, &x.TradeActionTrait)
	MapCustomToTradeActionTrait(ctx, &x.TradeActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRentActionTrait(ctx, &x.RentActionTrait)
	MapCustomToRentActionTrait(ctx, &x.RentActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRentActionTrait(ctx jsonld.Context, x *schema.RentActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRentActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRentActionTrait(ctx jsonld.Context, x *schema.RentActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRentActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}