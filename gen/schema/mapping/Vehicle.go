package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVehicle strings.Replacer
var strconvVehicle strconv.NumError

var basicVehicleTraitMapping = map[string]func(*schema.VehicleTrait, []string){}
var customVehicleTraitMapping = map[string]func(*schema.VehicleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Vehicle", func(ctx jsonld.Context) (interface{}) {
		return NewVehicleFromContext(ctx)
	})






	basicVehicleTraitMapping["http://schema.org/knownVehicleDamages"] = func(x *schema.VehicleTrait, s []string) {
		x.KnownVehicleDamages = s[0]
	}




	basicVehicleTraitMapping["http://schema.org/vehicleIdentificationNumber"] = func(x *schema.VehicleTrait, s []string) {
		x.VehicleIdentificationNumber = s[0]
	}


	basicVehicleTraitMapping["http://schema.org/vehicleConfiguration"] = func(x *schema.VehicleTrait, s []string) {
		x.VehicleConfiguration = s[0]
	}




	basicVehicleTraitMapping["http://schema.org/vehicleInteriorType"] = func(x *schema.VehicleTrait, s []string) {
		x.VehicleInteriorType = s[0]
	}


	basicVehicleTraitMapping["http://schema.org/vehicleInteriorColor"] = func(x *schema.VehicleTrait, s []string) {
		x.VehicleInteriorColor = s[0]
	}























	basicVehicleTraitMapping["http://schema.org/emissionsCO2"] = func(x *schema.VehicleTrait, s []string) {
		var err error
		x.EmissionsCO2, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}





	customVehicleTraitMapping["http://schema.org/fuelType"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FuelType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FuelType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FuelType = s
		}
	}

	customVehicleTraitMapping["http://schema.org/productionDate"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.ProductionDate = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/purchaseDate"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.PurchaseDate = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/driveWheelConfiguration"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DriveWheelConfiguration, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DriveWheelConfiguration = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DriveWheelConfiguration = s
		}
	}

	customVehicleTraitMapping["http://schema.org/numberOfAirbags"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfAirbags, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfAirbags = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfAirbags = s
		}
	}

	customVehicleTraitMapping["http://schema.org/vehicleTransmission"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.VehicleTransmission, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.VehicleTransmission = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.VehicleTransmission = s
		}
	}

	customVehicleTraitMapping["http://schema.org/bodyType"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.BodyType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.BodyType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.BodyType = s
		}
	}

	customVehicleTraitMapping["http://schema.org/vehicleSpecialUsage"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.VehicleSpecialUsage, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.VehicleSpecialUsage = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.VehicleSpecialUsage = s
		}
	}

	customVehicleTraitMapping["http://schema.org/meetsEmissionStandard"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MeetsEmissionStandard, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MeetsEmissionStandard = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MeetsEmissionStandard = s
		}
	}

	customVehicleTraitMapping["http://schema.org/numberOfDoors"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfDoors, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfDoors = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfDoors = s
		}
	}

	customVehicleTraitMapping["http://schema.org/numberOfAxles"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfAxles, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfAxles = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfAxles = s
		}
	}

	customVehicleTraitMapping["http://schema.org/fuelCapacity"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.FuelCapacity = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/payload"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Payload = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/wheelbase"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Wheelbase = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/speed"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Speed = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/tongueWeight"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.TongueWeight = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/accelerationTime"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.AccelerationTime = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/mileageFromOdometer"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.MileageFromOdometer = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/fuelConsumption"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.FuelConsumption = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/trailerWeight"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.TrailerWeight = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/fuelEfficiency"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.FuelEfficiency = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/cargoVolume"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.CargoVolume = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/numberOfPreviousOwners"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfPreviousOwners, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfPreviousOwners = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfPreviousOwners = s
		}
	}

	customVehicleTraitMapping["http://schema.org/numberOfForwardGears"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfForwardGears, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfForwardGears = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfForwardGears = s
		}
	}

	customVehicleTraitMapping["http://schema.org/vehicleSeatingCapacity"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.VehicleSeatingCapacity, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.VehicleSeatingCapacity = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.VehicleSeatingCapacity = s
		}
	}

	customVehicleTraitMapping["http://schema.org/seatingCapacity"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SeatingCapacity, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SeatingCapacity = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SeatingCapacity = s
		}
	}

	customVehicleTraitMapping["http://schema.org/weightTotal"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.WeightTotal = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/vehicleEngine"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.EngineSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewEngineSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEngineSpecification()
			y.Id = s
		}

		x.VehicleEngine = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/dateVehicleFirstRegistered"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.DateVehicleFirstRegistered = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/modelDate"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.ModelDate = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/steeringPosition"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.SteeringPositionValue
		if strings.HasPrefix(s, "_:") {
			y = NewSteeringPositionValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSteeringPositionValue()
			y.Id = s
		}

		x.SteeringPosition = &y

		return
	}

	customVehicleTraitMapping["http://schema.org/vehicleModelDate"] = func(x *schema.VehicleTrait, ctx jsonld.Context, s string){
		var y schema.Date
		if strings.HasPrefix(s, "_:") {
			y = NewDateFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDate()
			y.Id = s
		}

		x.VehicleModelDate = &y

		return
	}
}

func NewVehicleFromContext(ctx jsonld.Context) (x schema.Vehicle) {
	x.Type = "http://schema.org/Vehicle"
	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVehicleTrait(ctx, &x.VehicleTrait)
	MapCustomToVehicleTrait(ctx, &x.VehicleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVehicleTrait(ctx jsonld.Context, x *schema.VehicleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVehicleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVehicleTrait(ctx jsonld.Context, x *schema.VehicleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVehicleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}