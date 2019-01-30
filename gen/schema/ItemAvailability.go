package schema

type ItemAvailabilityTrait struct {

}

type ItemAvailability struct {
	MetaTrait
	ItemAvailabilityTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewItemAvailability() (x ItemAvailability) {
	x.Type = "http://schema.org/ItemAvailability"

	return
}
