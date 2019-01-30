package schema

type BridgeTrait struct {

}

type Bridge struct {
	MetaTrait
	BridgeTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewBridge() (x Bridge) {
	x.Type = "http://schema.org/Bridge"

	return
}
