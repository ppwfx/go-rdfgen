package schema

type LibraryTrait struct {

}

type Library struct {
	MetaTrait
	LibraryTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewLibrary() (x Library) {
	x.Type = "http://schema.org/Library"

	return
}
