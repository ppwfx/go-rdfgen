package schema

type MobilePhoneStoreTrait struct {

}

type MobilePhoneStore struct {
	MetaTrait
	MobilePhoneStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMobilePhoneStore() (x MobilePhoneStore) {
	x.Type = "http://schema.org/MobilePhoneStore"

	return
}
