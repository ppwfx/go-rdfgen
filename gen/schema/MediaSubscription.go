package schema

type MediaSubscriptionTrait struct {

	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty" jsonld:"http://schema.org/expectsAcceptanceOf"`

	// The Organization responsible for authenticating the user's subscription. For example, many media apps require a cable/satellite provider to authenticate your subscription before playing media.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Authenticator *Organization `json:"authenticator,omitempty" jsonld:"http://schema.org/authenticator"`

}

type MediaSubscription struct {
	MetaTrait
	MediaSubscriptionTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMediaSubscription() (x MediaSubscription) {
	x.Type = "http://schema.org/MediaSubscription"

	return
}
