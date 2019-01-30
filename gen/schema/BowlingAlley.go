package schema

type BowlingAlleyTrait struct {

}

type BowlingAlley struct {
	MetaTrait
	BowlingAlleyTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBowlingAlley() (x BowlingAlley) {
	x.Type = "http://schema.org/BowlingAlley"

	return
}
