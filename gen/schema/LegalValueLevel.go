package schema

type LegalValueLevelTrait struct {

}

type LegalValueLevel struct {
	MetaTrait
	LegalValueLevelTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLegalValueLevel() (x LegalValueLevel) {
	x.Type = "http://schema.org/LegalValueLevel"

	return
}
