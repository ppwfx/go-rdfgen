package schema

type CommunicateActionTrait struct {

	// The subject matter of the content.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	About *Thing `json:"about,omitempty" jsonld:"http://schema.org/about"`

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

	// A sub property of participant. The participant who is at the receiving end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Audience
	// - http://schema.org/Organization
	//
	Recipient interface{} `json:"recipient,omitempty" jsonld:"http://schema.org/recipient"`

}

type CommunicateAction struct {
	MetaTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCommunicateAction() (x CommunicateAction) {
	x.Type = "http://schema.org/CommunicateAction"

	return
}
