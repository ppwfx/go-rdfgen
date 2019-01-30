package schema

type DurationTrait struct {

}

type Duration struct {
	MetaTrait
	DurationTrait
	QuantityTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDuration() (x Duration) {
	x.Type = "http://schema.org/Duration"

	return
}
