package schema

type LibrarySystemTrait struct {

}

type LibrarySystem struct {
	MetaTrait
	LibrarySystemTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewLibrarySystem() (x LibrarySystem) {
	x.Type = "http://schema.org/LibrarySystem"

	return
}
