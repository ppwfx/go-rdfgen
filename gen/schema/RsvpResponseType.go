package schema

type RsvpResponseTypeTrait struct {

}

type RsvpResponseType struct {
	MetaTrait
	RsvpResponseTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRsvpResponseType() (x RsvpResponseType) {
	x.Type = "http://schema.org/RsvpResponseType"

	return
}
