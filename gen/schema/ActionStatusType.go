package schema

type ActionStatusTypeTrait struct {

}

type ActionStatusType struct {
	MetaTrait
	ActionStatusTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewActionStatusType() (x ActionStatusType) {
	x.Type = "http://schema.org/ActionStatusType"

	return
}
