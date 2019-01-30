package schema

type DiscoverActionTrait struct {

}

type DiscoverAction struct {
	MetaTrait
	DiscoverActionTrait
	FindActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDiscoverAction() (x DiscoverAction) {
	x.Type = "http://schema.org/DiscoverAction"

	return
}
