package schema

type DepartActionTrait struct {

}

type DepartAction struct {
	MetaTrait
	DepartActionTrait
	MoveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDepartAction() (x DepartAction) {
	x.Type = "http://schema.org/DepartAction"

	return
}
