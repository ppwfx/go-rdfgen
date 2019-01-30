package schema

type HousePainterTrait struct {

}

type HousePainter struct {
	MetaTrait
	HousePainterTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHousePainter() (x HousePainter) {
	x.Type = "http://schema.org/HousePainter"

	return
}
