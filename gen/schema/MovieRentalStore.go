package schema

type MovieRentalStoreTrait struct {

}

type MovieRentalStore struct {
	MetaTrait
	MovieRentalStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMovieRentalStore() (x MovieRentalStore) {
	x.Type = "http://schema.org/MovieRentalStore"

	return
}
