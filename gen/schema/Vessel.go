package schema

type VesselTrait struct {

}

type Vessel struct {
	MetaTrait
	VesselTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewVessel() (x Vessel) {
	x.Type = "http://schema.org/Vessel"

	return
}
