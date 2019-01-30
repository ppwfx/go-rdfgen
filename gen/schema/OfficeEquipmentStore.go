package schema

type OfficeEquipmentStoreTrait struct {

}

type OfficeEquipmentStore struct {
	MetaTrait
	OfficeEquipmentStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewOfficeEquipmentStore() (x OfficeEquipmentStore) {
	x.Type = "http://schema.org/OfficeEquipmentStore"

	return
}
