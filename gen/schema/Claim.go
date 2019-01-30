package schema

type ClaimTrait struct {

	// Indicates an occurence of a <a class=\"localLink\" href=\"http://schema.org/Claim\">Claim</a> in some <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a>.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	Appearance *CreativeWork `json:"appearance,omitempty" jsonld:"http://schema.org/appearance"`

	// Indicates the first known occurence of a <a class=\"localLink\" href=\"http://schema.org/Claim\">Claim</a> in some <a class=\"localLink\" href=\"http://schema.org/CreativeWork\">CreativeWork</a>.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	FirstAppearance *CreativeWork `json:"firstAppearance,omitempty" jsonld:"http://schema.org/firstAppearance"`

}

type Claim struct {
	MetaTrait
	ClaimTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewClaim() (x Claim) {
	x.Type = "http://schema.org/Claim"

	return
}
