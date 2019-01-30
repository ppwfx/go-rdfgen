package schema

type FoodEstablishmentTrait struct {

	// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	//
	// RangeIncludes:
	// - http://schema.org/Rating
	//
	StarRating *Rating `json:"starRating,omitempty" jsonld:"http://schema.org/starRating"`

	// Either the actual menu as a structured representation, as text, or a URL of the menu.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/Menu
	//
	HasMenu interface{} `json:"hasMenu,omitempty" jsonld:"http://schema.org/hasMenu"`

	// The cuisine of the restaurant.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ServesCuisine string `json:"servesCuisine,omitempty" jsonld:"http://schema.org/servesCuisine"`

	// Indicates whether a FoodEstablishment accepts reservations. Values can be Boolean, an URL at which reservations can be made or (for backwards compatibility) the strings <code>Yes</code> or <code>No</code>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/Boolean
	//
	AcceptsReservations interface{} `json:"acceptsReservations,omitempty" jsonld:"http://schema.org/acceptsReservations"`

	// Either the actual menu as a structured representation, as text, or a URL of the menu.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/Menu
	//
	Menu interface{} `json:"menu,omitempty" jsonld:"http://schema.org/menu"`

}

type FoodEstablishment struct {
	MetaTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewFoodEstablishment() (x FoodEstablishment) {
	x.Type = "http://schema.org/FoodEstablishment"

	return
}
