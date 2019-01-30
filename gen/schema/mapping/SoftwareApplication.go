package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSoftwareApplication strings.Replacer
var strconvSoftwareApplication strconv.NumError

var basicSoftwareApplicationTraitMapping = map[string]func(*schema.SoftwareApplicationTrait, []string){}
var customSoftwareApplicationTraitMapping = map[string]func(*schema.SoftwareApplicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SoftwareApplication", func(ctx jsonld.Context) (interface{}) {
		return NewSoftwareApplicationFromContext(ctx)
	})





	basicSoftwareApplicationTraitMapping["http://schema.org/applicationSuite"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.ApplicationSuite = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/availableOnDevice"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.AvailableOnDevice = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/countriesNotSupported"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.CountriesNotSupported = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/countriesSupported"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.CountriesSupported = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/device"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.Device = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/downloadUrl"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.DownloadUrl = s[0]
	}



	basicSoftwareApplicationTraitMapping["http://schema.org/fileSize"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.FileSize = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/installUrl"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.InstallUrl = s[0]
	}



	basicSoftwareApplicationTraitMapping["http://schema.org/operatingSystem"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.OperatingSystem = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/permissions"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.Permissions = s[0]
	}


	basicSoftwareApplicationTraitMapping["http://schema.org/processorRequirements"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.ProcessorRequirements = s[0]
	}







	basicSoftwareApplicationTraitMapping["http://schema.org/softwareVersion"] = func(x *schema.SoftwareApplicationTrait, s []string) {
		x.SoftwareVersion = s[0]
	}




	customSoftwareApplicationTraitMapping["http://schema.org/screenshot"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Screenshot, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Screenshot = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Screenshot = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/applicationCategory"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ApplicationCategory, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ApplicationCategory = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ApplicationCategory = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/applicationSubCategory"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ApplicationSubCategory, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ApplicationSubCategory = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ApplicationSubCategory = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/featureList"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FeatureList, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FeatureList = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FeatureList = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/memoryRequirements"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MemoryRequirements, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MemoryRequirements = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MemoryRequirements = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/releaseNotes"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ReleaseNotes, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ReleaseNotes = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ReleaseNotes = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/requirements"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Requirements, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Requirements = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Requirements = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/softwareAddOn"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		var y schema.SoftwareApplication
		if strings.HasPrefix(s, "_:") {
			y = NewSoftwareApplicationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewSoftwareApplication()
			y.Id = s
		}

		x.SoftwareAddOn = &y

		return
	}

	customSoftwareApplicationTraitMapping["http://schema.org/softwareHelp"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		var y schema.CreativeWork
		if strings.HasPrefix(s, "_:") {
			y = NewCreativeWorkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewCreativeWork()
			y.Id = s
		}

		x.SoftwareHelp = &y

		return
	}

	customSoftwareApplicationTraitMapping["http://schema.org/softwareRequirements"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.SoftwareRequirements, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.SoftwareRequirements = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.SoftwareRequirements = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/storageRequirements"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.StorageRequirements, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.StorageRequirements = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.StorageRequirements = s
		}
	}

	customSoftwareApplicationTraitMapping["http://schema.org/supportingData"] = func(x *schema.SoftwareApplicationTrait, ctx jsonld.Context, s string){
		var y schema.DataFeed
		if strings.HasPrefix(s, "_:") {
			y = NewDataFeedFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDataFeed()
			y.Id = s
		}

		x.SupportingData = &y

		return
	}
}

func NewSoftwareApplicationFromContext(ctx jsonld.Context) (x schema.SoftwareApplication) {
	x.Type = "http://schema.org/SoftwareApplication"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)
	MapCustomToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSoftwareApplicationTrait(ctx jsonld.Context, x *schema.SoftwareApplicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSoftwareApplicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSoftwareApplicationTrait(ctx jsonld.Context, x *schema.SoftwareApplicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSoftwareApplicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}