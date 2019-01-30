package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsVoteAction strings.Replacer
var strconvVoteAction strconv.NumError

var basicVoteActionTraitMapping = map[string]func(*schema.VoteActionTrait, []string){}
var customVoteActionTraitMapping = map[string]func(*schema.VoteActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/VoteAction", func(ctx jsonld.Context) (interface{}) {
		return NewVoteActionFromContext(ctx)
	})



	customVoteActionTraitMapping["http://schema.org/candidate"] = func(x *schema.VoteActionTrait, ctx jsonld.Context, s string){
		var y schema.Person
		if strings.HasPrefix(s, "_:") {
			y = NewPersonFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPerson()
			y.Id = s
		}

		x.Candidate = &y

		return
	}
}

func NewVoteActionFromContext(ctx jsonld.Context) (x schema.VoteAction) {
	x.Type = "http://schema.org/VoteAction"
	MapBasicToChooseActionTrait(ctx, &x.ChooseActionTrait)
	MapCustomToChooseActionTrait(ctx, &x.ChooseActionTrait)

	MapBasicToAssessActionTrait(ctx, &x.AssessActionTrait)
	MapCustomToAssessActionTrait(ctx, &x.AssessActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToVoteActionTrait(ctx, &x.VoteActionTrait)
	MapCustomToVoteActionTrait(ctx, &x.VoteActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToVoteActionTrait(ctx jsonld.Context, x *schema.VoteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicVoteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToVoteActionTrait(ctx jsonld.Context, x *schema.VoteActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customVoteActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}