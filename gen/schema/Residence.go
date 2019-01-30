package schema

type ResidenceTrait struct {

}

type Residence struct {
	MetaTrait
	ResidenceTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewResidence() (x Residence) {
	x.Type = "http://schema.org/Residence"

	return
}
