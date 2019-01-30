package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDriveWheelConfigurationValue strings.Replacer
var strconvDriveWheelConfigurationValue strconv.NumError

var basicDriveWheelConfigurationValueTraitMapping = map[string]func(*schema.DriveWheelConfigurationValueTrait, []string){}
var customDriveWheelConfigurationValueTraitMapping = map[string]func(*schema.DriveWheelConfigurationValueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DriveWheelConfigurationValue", func(ctx jsonld.Context) (interface{}) {
		return NewDriveWheelConfigurationValueFromContext(ctx)
	})

}

func NewDriveWheelConfigurationValueFromContext(ctx jsonld.Context) (x schema.DriveWheelConfigurationValue) {
	x.Type = "http://schema.org/DriveWheelConfigurationValue"
	MapBasicToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)
	MapCustomToQualitativeValueTrait(ctx, &x.QualitativeValueTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDriveWheelConfigurationValueTrait(ctx, &x.DriveWheelConfigurationValueTrait)
	MapCustomToDriveWheelConfigurationValueTrait(ctx, &x.DriveWheelConfigurationValueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDriveWheelConfigurationValueTrait(ctx jsonld.Context, x *schema.DriveWheelConfigurationValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDriveWheelConfigurationValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDriveWheelConfigurationValueTrait(ctx jsonld.Context, x *schema.DriveWheelConfigurationValueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDriveWheelConfigurationValueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}