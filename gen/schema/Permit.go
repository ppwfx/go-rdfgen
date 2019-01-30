package schema

type PermitTrait struct {

	// The date when the item becomes valid.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidFrom *DateTime `json:"validFrom,omitempty" jsonld:"http://schema.org/validFrom"`

	// The organization issuing the ticket or permit.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	IssuedBy *Organization `json:"issuedBy,omitempty" jsonld:"http://schema.org/issuedBy"`

	// The service through with the permit was granted.
	//
	// RangeIncludes:
	// - http://schema.org/Service
	//
	IssuedThrough *Service `json:"issuedThrough,omitempty" jsonld:"http://schema.org/issuedThrough"`

	// The target audience for this permit.
	//
	// RangeIncludes:
	// - http://schema.org/Audience
	//
	PermitAudience *Audience `json:"permitAudience,omitempty" jsonld:"http://schema.org/permitAudience"`

	// The time validity of the permit.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	ValidFor *Duration `json:"validFor,omitempty" jsonld:"http://schema.org/validFor"`

	// The geographic area where the permit is valid.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	ValidIn *AdministrativeArea `json:"validIn,omitempty" jsonld:"http://schema.org/validIn"`

	// The date when the item is no longer valid.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	ValidUntil *Date `json:"validUntil,omitempty" jsonld:"http://schema.org/validUntil"`

}

type Permit struct {
	MetaTrait
	PermitTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPermit() (x Permit) {
	x.Type = "http://schema.org/Permit"

	return
}
