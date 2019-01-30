package schema

type WriteActionTrait struct {

	// The language of the content or performance or used in an action. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/availableLanguage\">availableLanguage</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Language
	// - http://schema.org/Text
	//
	InLanguage interface{} `json:"inLanguage,omitempty" jsonld:"http://schema.org/inLanguage"`

	// A sub property of instrument. The language used on this action.
	//
	// RangeIncludes:
	// - http://schema.org/Language
	//
	Language *Language `json:"language,omitempty" jsonld:"http://schema.org/language"`

}

type WriteAction struct {
	MetaTrait
	WriteActionTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewWriteAction() (x WriteAction) {
	x.Type = "http://schema.org/WriteAction"

	return
}
