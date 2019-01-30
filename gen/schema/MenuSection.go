package schema

type MenuSectionTrait struct {

	// A food or drink item contained in a menu or menu section.
	//
	// RangeIncludes:
	// - http://schema.org/MenuItem
	//
	HasMenuItem *MenuItem `json:"hasMenuItem,omitempty" jsonld:"http://schema.org/hasMenuItem"`

	// A subgrouping of the menu (by dishes, course, serving time period, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/MenuSection
	//
	HasMenuSection *MenuSection `json:"hasMenuSection,omitempty" jsonld:"http://schema.org/hasMenuSection"`

}

type MenuSection struct {
	MetaTrait
	MenuSectionTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMenuSection() (x MenuSection) {
	x.Type = "http://schema.org/MenuSection"

	return
}
