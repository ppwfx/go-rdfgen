package schema

type LegalForceStatusTrait struct {

}

type LegalForceStatus struct {
	MetaTrait
	LegalForceStatusTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLegalForceStatus() (x LegalForceStatus) {
	x.Type = "http://schema.org/LegalForceStatus"

	return
}
