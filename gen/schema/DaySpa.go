package schema

type DaySpaTrait struct {

}

type DaySpa struct {
	MetaTrait
	DaySpaTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewDaySpa() (x DaySpa) {
	x.Type = "http://schema.org/DaySpa"

	return
}
