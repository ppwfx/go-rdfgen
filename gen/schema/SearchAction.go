package schema

type SearchActionTrait struct {

	// A sub property of instrument. The query used on this action.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Query string `json:"query,omitempty" jsonld:"http://schema.org/query"`

}

type SearchAction struct {
	MetaTrait
	SearchActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewSearchAction() (x SearchAction) {
	x.Type = "http://schema.org/SearchAction"

	return
}
