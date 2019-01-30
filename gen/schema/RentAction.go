package schema

type RentActionTrait struct {

	// A sub property of participant. The owner of the real estate property.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Landlord interface{} `json:"landlord,omitempty" jsonld:"http://schema.org/landlord"`

	// A sub property of participant. The real estate agent involved in the action.
	//
	// RangeIncludes:
	// - http://schema.org/RealEstateAgent
	//
	RealEstateAgent *RealEstateAgent `json:"realEstateAgent,omitempty" jsonld:"http://schema.org/realEstateAgent"`

}

type RentAction struct {
	MetaTrait
	RentActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewRentAction() (x RentAction) {
	x.Type = "http://schema.org/RentAction"

	return
}
