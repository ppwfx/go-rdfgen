package schema

type BarOrPubTrait struct {

}

type BarOrPub struct {
	MetaTrait
	BarOrPubTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBarOrPub() (x BarOrPub) {
	x.Type = "http://schema.org/BarOrPub"

	return
}
