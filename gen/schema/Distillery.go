package schema

type DistilleryTrait struct {

}

type Distillery struct {
	MetaTrait
	DistilleryTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewDistillery() (x Distillery) {
	x.Type = "http://schema.org/Distillery"

	return
}
