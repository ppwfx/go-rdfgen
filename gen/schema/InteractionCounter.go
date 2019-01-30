package schema

type InteractionCounterTrait struct {

	// The WebSite or SoftwareApplication where the interactions took place.
	//
	// RangeIncludes:
	// - http://schema.org/SoftwareApplication
	// - http://schema.org/WebSite
	//
	InteractionService interface{} `json:"interactionService,omitempty" jsonld:"http://schema.org/interactionService"`

	// The Action representing the type of interaction. For up votes, +1s, etc. use <a class=\"localLink\" href=\"http://schema.org/LikeAction\">LikeAction</a>. For down votes use <a class=\"localLink\" href=\"http://schema.org/DislikeAction\">DislikeAction</a>. Otherwise, use the most specific Action.
	//
	// RangeIncludes:
	// - http://schema.org/Action
	//
	InteractionType *Action `json:"interactionType,omitempty" jsonld:"http://schema.org/interactionType"`

	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	UserInteractionCount *Integer `json:"userInteractionCount,omitempty" jsonld:"http://schema.org/userInteractionCount"`

}

type InteractionCounter struct {
	MetaTrait
	InteractionCounterTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewInteractionCounter() (x InteractionCounter) {
	x.Type = "http://schema.org/InteractionCounter"

	return
}
