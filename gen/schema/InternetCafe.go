package schema

type InternetCafeTrait struct {

}

type InternetCafe struct {
	MetaTrait
	InternetCafeTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewInternetCafe() (x InternetCafe) {
	x.Type = "http://schema.org/InternetCafe"

	return
}
