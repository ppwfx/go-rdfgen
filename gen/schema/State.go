package schema

type StateTrait struct {

}

type State struct {
	MetaTrait
	StateTrait
	AdministrativeAreaTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewState() (x State) {
	x.Type = "http://schema.org/State"

	return
}
