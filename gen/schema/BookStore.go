package schema

type BookStoreTrait struct {

}

type BookStore struct {
	MetaTrait
	BookStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBookStore() (x BookStore) {
	x.Type = "http://schema.org/BookStore"

	return
}
