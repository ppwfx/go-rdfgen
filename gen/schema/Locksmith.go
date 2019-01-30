package schema

type LocksmithTrait struct {

}

type Locksmith struct {
	MetaTrait
	LocksmithTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewLocksmith() (x Locksmith) {
	x.Type = "http://schema.org/Locksmith"

	return
}
