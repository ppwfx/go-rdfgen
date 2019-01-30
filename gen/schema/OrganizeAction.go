package schema

type OrganizeActionTrait struct {

}

type OrganizeAction struct {
	MetaTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewOrganizeAction() (x OrganizeAction) {
	x.Type = "http://schema.org/OrganizeAction"

	return
}
