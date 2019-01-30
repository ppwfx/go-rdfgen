package schema

type StadiumOrArenaTrait struct {

}

type StadiumOrArena struct {
	MetaTrait
	StadiumOrArenaTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	OrganizationTrait
	AdditionalTrait
}

func NewStadiumOrArena() (x StadiumOrArena) {
	x.Type = "http://schema.org/StadiumOrArena"

	return
}
