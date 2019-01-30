package schema

type TravelActionTrait struct {

	// The distance travelled, e.g. exercising or travelling.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	//
	Distance *Distance `json:"distance,omitempty" jsonld:"http://schema.org/distance"`

}

type TravelAction struct {
	MetaTrait
	TravelActionTrait
	MoveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewTravelAction() (x TravelAction) {
	x.Type = "http://schema.org/TravelAction"

	return
}
