package schema

type ContactPointOptionTrait struct {

}

type ContactPointOption struct {
	MetaTrait
	ContactPointOptionTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewContactPointOption() (x ContactPointOption) {
	x.Type = "http://schema.org/ContactPointOption"

	return
}
