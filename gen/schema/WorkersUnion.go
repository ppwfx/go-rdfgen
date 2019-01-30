package schema

type WorkersUnionTrait struct {

}

type WorkersUnion struct {
	MetaTrait
	WorkersUnionTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewWorkersUnion() (x WorkersUnion) {
	x.Type = "http://schema.org/WorkersUnion"

	return
}
