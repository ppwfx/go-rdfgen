package schema

type PerformingArtsTheaterTrait struct {

}

type PerformingArtsTheater struct {
	MetaTrait
	PerformingArtsTheaterTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPerformingArtsTheater() (x PerformingArtsTheater) {
	x.Type = "http://schema.org/PerformingArtsTheater"

	return
}
