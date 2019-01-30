package schema

type WineryTrait struct {

}

type Winery struct {
	MetaTrait
	WineryTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewWinery() (x Winery) {
	x.Type = "http://schema.org/Winery"

	return
}
