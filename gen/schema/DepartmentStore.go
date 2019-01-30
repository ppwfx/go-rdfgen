package schema

type DepartmentStoreTrait struct {

}

type DepartmentStore struct {
	MetaTrait
	DepartmentStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewDepartmentStore() (x DepartmentStore) {
	x.Type = "http://schema.org/DepartmentStore"

	return
}
